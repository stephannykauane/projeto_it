package services

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stephannykauane/projeto_it/backend/calculator"
	datab "github.com/stephannykauane/projeto_it/backend/db"
	"github.com/stephannykauane/projeto_it/backend/middleware"
	"github.com/stephannykauane/projeto_it/backend/types"
	"github.com/xuri/excelize/v2"
)

func SaveExcel(w http.ResponseWriter, r *http.Request) {
	var metodo types.Metodo
	var values types.AnaliseRequest

	fmt.Println("Chegamos aqui")

	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("passei por aqui")

	if err := json.Unmarshal(jsonData, &metodo); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar JSON para Metodo: %v", err), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(jsonData, &values); err != nil {
		http.Error(w, "Erro ao deserializar JSON para AnaliseRequest", http.StatusBadRequest)
		return
	}

	fmt.Println("metodo: ", metodo)
	fmt.Println("valores: ", values)

	templateFile := ""

	switch metodo.MetodoID {
	case 1:
		templateFile = "/home/stephanny/projeto_it/backend/cmd/RecomendaçãoCalagem.xlsx"
	case 2:
		templateFile = "/home/stephanny/projeto_it/backend/cmd/RecomendaçãoDeCalagem.xlsx"
	default:
		http.Error(w, "MetodoID inválido", http.StatusBadRequest)
		return
	}

	f, err := excelize.OpenFile(templateFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir arquivo template: %v", err), http.StatusInternalServerError)
		return
	}

	sheetName := "Recomendação de Calagem"
	if index, err := f.GetSheetIndex(sheetName); err != nil || index == -1 {
		f.NewSheet(sheetName)
	}

	resultado, err := calculator.Calculando(jsonData, metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao calcular:%v", err), http.StatusInternalServerError)

		return
	}
	fmt.Println("result:", resultado.Result.Float64)

	data := map[string]interface{}{}

	switch metodo.MetodoID {
	case 1:
		data = map[string]interface{}{
			"A11": values.SatA,
			"B10": values.SatD,
			"C10": values.Prnt,
			"D11": values.Ctc,
			"E11": values.Resultado,
		}
	case 2:
		data = map[string]interface{}{
			"A10": values.Argila,
			"B10": values.Magnesio,
			"C10": values.Calcio,
			"D10": values.Aluminio,
			"E10": values.Prnt,
			"F10": values.Ctc,
			"G10": values.Resultado,
		}
	}

	fmt.Println("dados da planilha: ", data)

	for cell, value := range data {
		if err := f.SetCellValue(sheetName, cell, value); err != nil {
			fmt.Printf("Erro ao definir valor na célula %s: %v\n", cell, err)
		}
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	if err := f.Write(w); err != nil {
		http.Error(w, fmt.Sprintln("Erro ao salvar arquivo: %w", err), http.StatusInternalServerError)
		return
	}

}

func SaveAnalise(w http.ResponseWriter, r *http.Request) {

	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var userID int
	err := db.QueryRow(`SELECT id FROM "Usuario" WHERE email=$1`, email).Scan(&userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	var analiseRequest types.AnaliseRequest
	var metodo types.Metodo
	var areaRequest types.Area

	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler corpo da requisição: %v", err), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonData, &analiseRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar JSON: %v", err), http.StatusBadRequest)
		return
	} 


	var areaID int64

	areaQuery := `INSERT INTO "Area" ("id_usuario", "consultor", "propriedade", "area") VALUES ($1, $2, $3, $4) RETURNING id`
	
	err = db.QueryRow(areaQuery, userID, areaRequest.Consultor, areaRequest.Propriedade, areaRequest.Area).Scan(&areaID)
	
	if err != nil {
		fmt.Println("Erro ao inserir área:", err)

		return
	}
	
	fmt.Println("ID da área inserida:", areaID)
	

	query := `INSERT INTO "Analise" ("id_usuario", "magnesio", "aluminio", "calcio", "sat_desejada", "prnt", "ctc", "argila", "sat_atual", "teor_ca", "teor_mg", "caO", "mgO", "mg_desejada", "ca_desejada", "id_area") 
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING "id"`

	var analiseID int
	err = db.QueryRow(query, userID, analiseRequest.Magnesio, analiseRequest.Aluminio, analiseRequest.Calcio, analiseRequest.SatD, analiseRequest.Prnt, analiseRequest.Ctc, analiseRequest.Argila, analiseRequest.SatA, analiseRequest.TeorCa, analiseRequest.TeorMg, analiseRequest.CaO, analiseRequest.MgO, analiseRequest.MgDesejada, analiseRequest.CaDesejada, areaID).Scan(&analiseID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir análise: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(jsonData, &metodo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar JSON para Metodo: %v", err), http.StatusBadRequest)
		return
	}

	resultado, err := calculator.Calculando(jsonData, metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao calcular: %v", err), http.StatusInternalServerError)
		return
	}

	calculoQuery := `INSERT INTO "Calculo" ("id_analise", "resultado", "data_calculo", "id_metodo") 
						 VALUES ($1, $2, current_date, $3)`

	_, err = db.Exec(calculoQuery, analiseID, resultado.Result.Float64, metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir cálculo: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Àrea, análise e cálculo realizados com sucesso",
		"resultado": resultado.Result.Float64,
	})
}

func Hash(senha string) string {
	h := sha256.New()
	h.Write([]byte(senha))

	return hex.EncodeToString(h.Sum(nil))
}

func EfetuarLogin(w http.ResponseWriter, r *http.Request) {
	var jwtKey = os.Getenv("JWT_SECRET")

	var user types.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Senha == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var storedHash string
	db := datab.ConnectDB()
	defer db.Close()

	err = db.QueryRow(`SELECT senha FROM "Usuario" WHERE email=$1`, user.Email).Scan(&storedHash)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		return
	}

	if subtle.ConstantTimeCompare([]byte(storedHash), []byte(Hash(user.Senha))) != 1 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &types.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func AlterarDados(w http.ResponseWriter, r *http.Request) {
	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var userID int
	err := db.QueryRow(`SELECT id FROM "Usuario" WHERE email = $1`, email).Scan(&userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	var user types.Usuario
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.Nome == "" && user.Senha == "" {
		http.Error(w, "Nenhum dado para alterar foi fornecido", http.StatusBadRequest)
		return
	}

	query := `UPDATE "Usuario" SET `
	params := []interface{}{}
	paramIndex := 1

	if user.Nome != "" {
		query += fmt.Sprintf("nome = $%d", paramIndex)
		params = append(params, user.Nome)
		paramIndex++
	}

	if user.Senha != "" {
		if len(params) > 0 {
			query += ", "
		}
		hashed := Hash(user.Senha)
		query += fmt.Sprintf("senha = $%d", paramIndex)
		params = append(params, hashed)
		paramIndex++
	}

	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	params = append(params, userID)

	_, err = db.Exec(query, params...)
	if err != nil {
		http.Error(w, "Erro ao atualizar dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Dados atualizados com sucesso",
	})
}

func EfetuarSignUp(w http.ResponseWriter, r *http.Request) {

	var jwtKey = os.Getenv("JWT_SECRET")
	var user types.Usuario

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	hashed := Hash(user.Senha)
	db := datab.ConnectDB()
	defer db.Close()

	_, err := db.Exec(`INSERT INTO "Usuario"(email, senha, nome) VALUES ($1, $2, $3)`, user.Email, hashed, user.Nome)
	if err != nil {
		http.Error(w, "Erro ao executar query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &types.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func GetDadosUsuario(w http.ResponseWriter, r *http.Request) {
	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var user types.Usuario

	err := db.QueryRow(`SELECT nome, email FROM "Usuario" where email = $1`, email).Scan(&user.Nome, &user.Email)

	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func ListarCalculos(w http.ResponseWriter, r *http.Request) {

	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var userID int
	err := db.QueryRow(`SELECT id FROM "Usuario" WHERE email=$1`, email).Scan(&userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	query := `
		SELECT 
			c.resultado, 
			c.data_calculo, 
			c.id_metodo
		FROM "Calculo" c
		JOIN "Analise" a ON c.id_analise = a.id
		WHERE a.id_usuario = $1
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		http.Error(w, "Erro ao consultar banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var calculos []types.CalculoDetalhes
	for rows.Next() {
		var calculo types.CalculoDetalhes
		err := rows.Scan(
			&calculo.Resultado,
			&calculo.DataCalculo,
			&calculo.MetodoID,
		)
		if err != nil {
			http.Error(w, "Erro ao ler os dados: "+err.Error(), http.StatusInternalServerError)
			return
		}
		calculos = append(calculos, calculo)
	}

	w.Header().Set("Content-Type", "application/json")

	jsEnc := json.NewEncoder(w)
	err = jsEnc.Encode(calculos)
	if err != nil {
		fmt.Println(err)
	}

}

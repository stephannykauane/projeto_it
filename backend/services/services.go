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
	"path/filepath"
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

	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(jsonData, &metodo); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar JSON para Metodo: %v", err), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(jsonData, &values); err != nil {
		http.Error(w, "Erro ao deserializar JSON para AnaliseRequest", http.StatusBadRequest)
		return
	}

	basePath := os.Getenv("TEMPLATE_DIR")
	if basePath == "" {
		basePath = "./assets" 
	}

	templateFile := ""

	switch metodo.MetodoID {
	case 1:
		templateFile = filepath.Join(basePath, "RecomendaçãoSaturacaoBases.xlsx")
	case 2:
		templateFile = filepath.Join(basePath, "RecomendaçãoSaturacaoMagnesio.xlsx") 
	case 3:
		templateFile = filepath.Join(basePath, "RecomendaçãoSaturacaoCalcio.xlsx")
	case 4:
		templateFile = filepath.Join(basePath, "RecomendaçãoAluminio.xlsx")
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
			"B3":  values.Propriedade,
			"B4":  values.Area,
			"B5":  values.Consultor,
			"A11": values.SatA,
			"B11": values.SatD,
			"C11": values.Prnt,
			"D11": values.Ctc,
			"E11": values.Resultado,
		}
	case 2:
		data = map[string]interface{}{
			"B3":  values.Propriedade,
			"B4":  values.Area,
			"B5":  values.Consultor,
			"A10": values.Argila,
			"B10": values.Magnesio,
			"C10": values.Calcio,
			"D10": values.Aluminio,
			"E10": values.Prnt,
			"F10": values.Ctc,
			"G10": values.Resultado,
		}
	case 3:
		data = map[string]interface{}{
			"B3":  values.Propriedade,
			"B4":  values.Area,
			"B5":  values.Consultor,
			"A11": values.Ctc,
			"B11": values.TeorCa,
			"C11": values.Prnt,
			"D11": values.CaO,
			"E11": values.CaDesejada,
			"F11": values.Resultado,
		}
	case 4:
		data = map[string]interface{}{
			"B3":  values.Propriedade,
			"B4":  values.Area,
			"B5":  values.Consultor,
			"A11": values.Ctc,
			"B11": values.TeorMg,
			"C11": values.Prnt,
			"D11": values.MgO,
			"E11": values.MgDesejada,
			"F11": values.Resultado,
		}
	}

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
	err := db.QueryRow(`SELECT id FROM usuario WHERE email=$1`, email).Scan(&userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	var analiseRequest types.AnaliseRequest
	var metodo types.Metodo

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
	areaQuery := `INSERT INTO area (id_usuario, consultor, propriedade, area) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(areaQuery, userID, analiseRequest.Consultor, analiseRequest.Propriedade, analiseRequest.Area).Scan(&areaID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir área: %v", err), http.StatusInternalServerError)
		return
	}

	var analiseID int
	query := `INSERT INTO analise (id_usuario, magnesio, aluminio, calcio, sat_desejada, prnt, ctc, argila, sat_atual, teor_ca, teor_mg, cao, mgo, mg_desejada, ca_desejada, id_area) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`
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

	calculoQuery := `INSERT INTO calculo (id_analise, resultado, data_calculo, id_metodo, sat_extra, relacao_ca_mg) 
					 VALUES ($1, $2, current_date, $3, $4, $5)`
	_, err = db.Exec(calculoQuery, analiseID, resultado.Result.Float64, metodo.MetodoID, resultado.SatExtra.Float64, resultado.RelacaoCaMg.Float64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao inserir cálculo: %v", err), http.StatusInternalServerError)
		return
	}

	var calculoDetalhado types.CalculoDetalhes
	queryDetalhe := `
		SELECT 
			c.resultado, 
			c.data_calculo, 
			c.id_metodo,
			c.sat_extra,
			c.relacao_ca_mg,
			a.magnesio,
			a.aluminio,
			a.calcio,
			a.sat_desejada,
			a.prnt,
			a.ctc,
			a.argila,
			a.sat_atual,
			a.teor_ca,
			a.cao,
			a.mgo,
			a.teor_mg,
			a.mg_desejada,
			a.ca_desejada,
			ar.consultor,
			ar.propriedade,
			ar.area
		FROM calculo c
		JOIN analise a ON c.id_analise = a.id
		JOIN area ar ON a.id_area = ar.id
		WHERE c.id_analise = $1
		ORDER BY c.data_calculo DESC
		LIMIT 1
	`
	err = db.QueryRow(queryDetalhe, analiseID).Scan(
		&calculoDetalhado.Resultado,
		&calculoDetalhado.DataCalculo,
		&calculoDetalhado.MetodoID,
		&calculoDetalhado.SatExtra,
		&calculoDetalhado.RelacaoCaMg,
		&calculoDetalhado.Magnesio,
		&calculoDetalhado.Aluminio,
		&calculoDetalhado.Calcio,
		&calculoDetalhado.SatDesejada,
		&calculoDetalhado.Prnt,
		&calculoDetalhado.Ctc,
		&calculoDetalhado.Argila,
		&calculoDetalhado.SatAtual,
		&calculoDetalhado.TeorCa,
		&calculoDetalhado.CaO,
		&calculoDetalhado.MgO,
		&calculoDetalhado.TeorMg,
		&calculoDetalhado.MgDesejada,
		&calculoDetalhado.CaDesejada,
		&calculoDetalhado.Consultor,
		&calculoDetalhado.Propriedade,
		&calculoDetalhado.Area,
	)
	if err != nil {
		http.Error(w, "Erro ao recuperar cálculo salvo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(calculoDetalhado)
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
	var userID int

	err = db.QueryRow(`SELECT id, senha FROM usuario WHERE email=$1`, user.Email).Scan(&userID, &storedHash)
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
			Subject:   fmt.Sprint(userID),
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
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func AlterarDados(w http.ResponseWriter, r *http.Request) {
	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var userID int
	err := db.QueryRow(`SELECT id FROM usuario WHERE email = $1`, email).Scan(&userID)
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

	query := `UPDATE usuario SET `
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

	var exists int
	err := db.QueryRow(`SELECT COUNT(*) FROM usuario WHERE email = $1`, user.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Erro ao verificar email", http.StatusInternalServerError)
		return
	}
	if exists > 0 {
		http.Error(w, "Email já cadastrado", http.StatusBadRequest)
		return
	}

	var userID int

	err = db.QueryRow(`INSERT INTO usuario(email, senha, nome) VALUES ($1, $2, $3) RETURNING id`, user.Email, hashed, user.Nome).Scan(&userID)
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
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func GetDadosUsuario(w http.ResponseWriter, r *http.Request) {
	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var user types.Usuario

	err := db.QueryRow(`SELECT nome, email FROM usuario where email = $1`, email).Scan(&user.Nome, &user.Email)

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
	err := db.QueryRow(`SELECT id FROM usuario WHERE email=$1`, email).Scan(&userID)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	limit := 10
	page := 1

	queryValues := r.URL.Query()
	if p := queryValues.Get("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := queryValues.Get("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var totalCount int
	countQuery := `
		SELECT COUNT(*) 
		FROM calculo c
		JOIN analise a ON c.id_analise = a.id
		WHERE a.id_usuario = $1
	`
	err = db.QueryRow(countQuery, userID).Scan(&totalCount)
	if err != nil {
		http.Error(w, "Erro ao contar cálculos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	query := `
		SELECT 
			c.resultado, 
			c.data_calculo, 
			c.id_metodo,
			a.magnesio,
			a.aluminio,
			a.calcio,
			a.sat_desejada,
			a.prnt,
			a.ctc,
			a.argila,
			a.sat_atual,
			a.teor_ca,
			a.cao,
			a.mgO,
			a.teor_mg,
			a.mg_desejada,
			a.ca_desejada,
			ar.consultor,
			ar.propriedade,
			ar.area
		FROM calculo c
		JOIN analise a ON c.id_analise = a.id
		JOIN area ar ON a.id_area = ar.id
		WHERE a.id_usuario = $1
		ORDER BY c.data_calculo DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := db.Query(query, userID, limit, offset)
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
			&calculo.Magnesio,
			&calculo.Aluminio,
			&calculo.Calcio,
			&calculo.SatDesejada,
			&calculo.Prnt,
			&calculo.Ctc,
			&calculo.Argila,
			&calculo.SatAtual,
			&calculo.TeorCa,
			&calculo.CaO,
			&calculo.MgO,
			&calculo.TeorMg,
			&calculo.MgDesejada,
			&calculo.CaDesejada,
			&calculo.Consultor,
			&calculo.Propriedade,
			&calculo.Area,
		)
		if err != nil {
			http.Error(w, "Erro ao ler os dados: "+err.Error(), http.StatusInternalServerError)
			return
		}
		calculos = append(calculos, calculo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"dados":       calculos,
		"totalCount":  totalCount,
		"totalPages":  int((totalCount + limit - 1) / limit),
		"currentPage": page,
	})
}

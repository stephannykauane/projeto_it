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


	templateFile := ""
	outputFile := ""

	switch metodo.MetodoID {
	case 1:
		templateFile = "RecomendaçãoCalagem.xlsx"
		outputFile = "RecomendaçãoCalagem.xlsx"
	case 2:
		templateFile = "RecomendaçãoDeCalagem.xlsx"
		outputFile = "RecomendaçãoDeCalagem.xlsx"
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
    
	resultado, satAtual, err := calculator.Calculando(jsonData, metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao calcular:%v", err), http.StatusInternalServerError)
		
		return
	}

	data := map[string]interface{}{}
	fmt.Println("result:", resultado);
	switch metodo.MetodoID {
	case 1:
		data = map[string]interface{}{
			"A10": values.Potassio,
			"B10": values.Magnesio,
			"C10": values.Calcio,
			"D10": values.SatD,
			"E10": satAtual,
			"F10": values.SatD,
			"G10": values.Ctc,
			"H10": resultado,
		}
	case 2:
		
		data = map[string]interface{}{
			"A10": values.Argila,
			"B10": values.Magnesio,
			"C10": values.Calcio,
			"D10": values.Prnt,
			"E10": values.Ctc,
			"F10": resultado,
		}
	}


	for cell, value := range data {
		if err := f.SetCellValue(sheetName, cell, value); err != nil {
			fmt.Printf("Erro ao definir valor na célula %s: %v\n", cell, err)
		}
	}


	if err := f.SaveAs(outputFile); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao salvar arquivo: %v", err), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Planilha gerada com sucesso: %s", outputFile)))
}


func SaveAnalise (w http.ResponseWriter, r *http.Request){

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
	
	
		query := `INSERT INTO "Analise" ("id_usuario", "potassio", "magnesio", "aluminio", "calcio", "sat_desejada", "prnt", "ctc", "argila") 
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING "id"`
	
		var analiseID int
		err = db.QueryRow(query, userID, analiseRequest.Potassio, analiseRequest.Magnesio, analiseRequest.Aluminio, analiseRequest.Calcio, analiseRequest.SatD, analiseRequest.Prnt, analiseRequest.Ctc, analiseRequest.Argila).Scan(&analiseID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao inserir análise: %v", err), http.StatusInternalServerError)
			return
		}
	
		err = json.Unmarshal(jsonData, &metodo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao decodificar JSON para Metodo: %v", err), http.StatusBadRequest)
			return
		}
	
		resultado, _, err := calculator.Calculando(jsonData, metodo.MetodoID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao calcular: %v", err), http.StatusInternalServerError)
			return
		}
	
		calculoQuery := `INSERT INTO "Calculo" ("id_analise", "resultado", "data_calculo", "id_metodo") 
						 VALUES ($1, $2, current_date, $3)`
		_, err = db.Exec(calculoQuery, analiseID, resultado, metodo.MetodoID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao inserir cálculo: %v", err), http.StatusInternalServerError)
			return
		}
	
		w.WriteHeader(http.StatusCreated)
}


func Hash (senha string) string{
	h := sha256.New()
	h.Write([]byte(senha))

	return hex.EncodeToString(h.Sum(nil))
}

func EfetuarLogin (w http.ResponseWriter, r *http.Request){
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
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func EfetuarSignUp (w http.ResponseWriter, r *http.Request){

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
        http.Error(w, "Erro ao executar query: " + err.Error(), http.StatusInternalServerError)
        return
    }
}

func GetDadosUsuario (w http.ResponseWriter, r *http.Request) {
	db := datab.ConnectDB()
	defer db.Close()

	email := middleware.GetUserEmailFromContext(r.Context())

	var user types.Usuario

	err:= db.QueryRow(`SELECT nome, email FROM "Usuario" where email = $1`, email).Scan(&user.Nome, &user.Email)

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
			c.id_metodo, 
			a.potassio, 
			a.magnesio, 
			a.calcio, 
			a.sat_desejada, 
			a.prnt, 
			a.ctc, 
			a.argila
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
			&calculo.Potassio,
			&calculo.Magnesio,
			&calculo.Calcio,
			&calculo.SatDesejada,
			&calculo.Prnt,
			&calculo.Ctc,
			&calculo.Argila,
		)
		if err != nil {
			http.Error(w, "Erro ao ler os dados: "+err.Error(), http.StatusInternalServerError)
			return
		}
		calculos = append(calculos, calculo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calculos)
}
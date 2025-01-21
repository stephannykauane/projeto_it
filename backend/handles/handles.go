package handles

import (
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	datab "github.com/stephannykauane/projeto_it/backend/db"
    "github.com/golang-jwt/jwt/v4"
	"github.com/stephannykauane/projeto_it/backend/calculator"
	"github.com/stephannykauane/projeto_it/backend/headers"
	"github.com/stephannykauane/projeto_it/backend/types"
	"fmt"
	"io"
    "crypto/subtle"
    "time"
	"encoding/hex"
	"crypto/sha256"
	

 
)
func Hash (senha string) string{
	h := sha256.New()
	h.Write([]byte(senha))

	return hex.EncodeToString(h.Sum(nil))
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
    headers.SetHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

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

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &types.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
    var jwtKey = []byte("801a28164703e91603002e1a64bd675f6788dc1adef7ce1ae407fa155a52ccd09ca3eb969398378b6cb04a7f2ae756fdbf24b0b340e548439ff0eba371430532") 

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}


func PostSignUp(w http.ResponseWriter, r *http.Request) {
    headers.SetHeaders(w)

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"error": "Method not allowed"}`))
        return
    }

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

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Usuário cadastrado com sucesso!"}`))
}


func CriarAnalise(w http.ResponseWriter, r *http.Request) {
	var analiseRequest types.AnaliseRequest
	var metodo types.Metodo

	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao ler o corpo da requisição: %v", err), http.StatusBadRequest)
		return
	}


	err = json.Unmarshal(jsonData, &analiseRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao decodificar JSON: %v", err), http.StatusBadRequest)
		return
	}


	db := datab.ConnectDB()
	defer db.Close()

	var usuarioExists bool

	queryUsuario := `SELECT EXISTS (SELECT 1 FROM "Usuario" WHERE "id" = $1)`

	err = db.QueryRow(queryUsuario, analiseRequest.UsuarioID).Scan(&usuarioExists)
	
	if err != nil {
        http.Error(w, fmt.Sprintf("Erro ao verificar usuário no banco: %v", err), http.StatusInternalServerError)
        return
    }

    if !usuarioExists {
        http.Error(w, "Usuário não encontrado", http.StatusBadRequest)
        return
    }
    fmt.Println("ID do usuário recebido:", analiseRequest.UsuarioID)

	query := `INSERT INTO "Analise" ("id_usuario", "potassio", "magnesio", "aluminio", "calcio", "sat_desejada", "prnt", "ctc", "argila") 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING "id"`

	var analiseID int
	err = db.QueryRow(query, analiseRequest.UsuarioID, analiseRequest.Potassio, analiseRequest.Magnesio, analiseRequest.Aluminio, analiseRequest.Calcio, analiseRequest.SatD, analiseRequest.Prnt, analiseRequest.Ctc, analiseRequest.Argila).Scan(&analiseID)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao inserir análise: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(jsonData, &metodo)
    if err != nil {
	http.Error(w, fmt.Sprintf("erro ao decodificar JSON para Metodo: %v", err), http.StatusBadRequest)
	
	return
    }
    
	resultado, err := calculator.Calculando(jsonData, metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao calcular: %v", err), http.StatusInternalServerError)
		return
	}
    
	calculoQuery := `INSERT INTO "Calculo" ("id_analise", "resultado", "data_calculo", "id_metodo") 
	                 VALUES ($1, $2, current_date, $3)`
	_, err = db.Exec(calculoQuery, analiseID, resultado, metodo.MetodoID)
	fmt.Println("ID metodo recebido:", metodo.MetodoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao inserir cálculo: %v", err), http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Análise e cálculo realizados com sucesso",
		"resultado": resultado,
	})
}
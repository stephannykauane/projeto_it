package api

import (
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	datab "github.com/stephannykauane/projeto_it/backend/db"
    "crypto/subtle"
    "time"
    "github.com/golang-jwt/jwt/v4"
    


)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var user Usuario
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

	if subtle.ConstantTimeCompare([]byte(storedHash), []byte(hash(user.Senha))) != 1 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

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
    SetHeaders(w)

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"error": "Method not allowed"}`))
        return
    }

    var user Usuario
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
        return
    }

    hashed := hash(user.Senha)
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

func PostAnalise(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	var analise calculator.Analise
	if err := json.NewDecoder(r.Body).Decode(&analise); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	db := datab.ConnectDB()
	defer db.Close()

	query := `
		INSERT INTO analise (metodo_id, calcio, magnesio, potassio, ctc, sat_d, argila, aluminio, prnt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	var insertedID int
	err := db.QueryRow(query, analise.MetodoID, analise.Calcio, analise.Magnesio, analise.Potassio,
		analise.Ctc, analise.SatD, analise.Argila, analise.Aluminio, analise.Prnt).Scan(&insertedID)
	if err != nil {
		http.Error(w, "Erro ao inserir análise: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resultado, err := calculator.CalcularComDados(analise)
	if err != nil {
		http.Error(w, "Erro ao realizar cálculo: "+err.Error(), http.StatusInternalServerError)
		return
	}


	response := map[string]interface{}{
		"analise_id": insertedID,
		"resultado":  resultado,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



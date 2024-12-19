package api

import (
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	datab "github.com/stephannykauane/projeto_it/backend/db"

)

func GetLogin(w http.ResponseWriter, r *http.Request) {
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
   

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return 
    }
    
	db := datab.ConnectDB()
    defer db.Close()




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

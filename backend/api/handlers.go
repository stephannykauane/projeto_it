package api

import (
	"encoding/json"
	"net/http"
	datab "github.com/stephannykauane/projeto_it/backend/db"
	_ "github.com/lib/pq"
	
)
func GetLogin(w http.ResponseWriter, r *http.Request){

}

func PostSignUp(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		var user Usuario
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return 
		}
     
		db := datab.ConnectDB()
		_, err = db.Exec("INSERT INTO Usuario (email, senha, nome) VALUES ($1, $2, $3)", user.Email, user.Senha, user.Nome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	} else{
		http.Error(w, "Invalid request method", http.StatusAccepted)
	}
}

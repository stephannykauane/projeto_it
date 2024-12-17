package api

import (
	"encoding/json"
	"net/http"
	
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
	}
}

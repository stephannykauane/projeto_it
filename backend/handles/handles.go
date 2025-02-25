package handles

import (
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"	
	"github.com/stephannykauane/projeto_it/backend/headers"
	"github.com/stephannykauane/projeto_it/backend/services"
)



func Login(w http.ResponseWriter, r *http.Request) {

    headers.SetHeaders(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	services.EfetuarLogin(w, r)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}


func SignUp(w http.ResponseWriter, r *http.Request) {

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

	services.EfetuarSignUp(w, r)

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "Usuário cadastrado com sucesso!"}`))
}


func Analise(w http.ResponseWriter, r *http.Request) {

    headers.SetHeaders(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	services.SaveAnalise(w, r)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Análise e cálculo realizados com sucesso",
	})
}

func Excel (w http.ResponseWriter, r *http.Request){
	
	headers.SetHeaders(w)

	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost{
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"error": "Method not allowed"}`))
        return
    }

	services.SaveExcel(w, r)

}

func ListaCalculos (w http.ResponseWriter, r *http.Request){
	headers.SetHeaders(w)
	services.ListarCalculos(w, r)
}


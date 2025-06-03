package handles

import (
	"net/http"
	"time"
	_ "github.com/lib/pq"
	"github.com/stephannykauane/projeto_it/backend/services"
	
)



func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	services.EfetuarLogin(w, r)
}


func SignUp(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"error": "Method not allowed"}`))
        return
    }

	services.EfetuarSignUp(w, r)

}


func Analise(w http.ResponseWriter, r *http.Request) {
	services.SaveAnalise(w, r)
}



func Excel(w http.ResponseWriter, r *http.Request){
	
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

func DadosUsuario (w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPatch{
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte(`{"error": "Method not allowed"}`))
        return
    }
	services.AlterarDados(w, r)
}



func ListaCalculos (w http.ResponseWriter, r *http.Request){
	services.ListarCalculos(w, r)
}

func PerfilUsuario (w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}

	services.GetDadosUsuario(w, r)
}

func Logout (w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: "", 
		Expires: time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Secure: false,
		SameSite: http.SameSiteStrictMode,
		Path: "/",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successfull"))
}


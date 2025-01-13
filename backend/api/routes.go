package api

import (
	"net/http"
	"github.com/stephannykauane/projeto_it/backend/handles"

)

func Routes() {
	http.HandleFunc("/signup", PostSignUp)
	http.HandleFunc("/login", GetLogin)
	http.HandleFunc("/calculator", CriarAnalise)
	
}

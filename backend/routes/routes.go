package api

import (
	"net/http"
    "github.com/stephannykauane/projeto_it/backend/handles"
)

func Routes() {
	http.HandleFunc("/signup", handles.PostSignUp)
	http.HandleFunc("/login", handles.GetLogin)
	http.HandleFunc("/calculator", handles.CriarAnalise)

}

package routes

import (
	"net/http"
    "github.com/stephannykauane/projeto_it/backend/handles"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
    http.HandleFunc("/signup", handles.SignUp)
	http.HandleFunc("/login", handles.Login)
	http.HandleFunc("/calculator", handles.Analise)
	http.HandleFunc("/excel", handles.Excel)
	http.HandleFunc("/historico", handles.ListaCalculos)

}

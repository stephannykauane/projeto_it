package routes

import (
	"net/http"

	"github.com/stephannykauane/projeto_it/backend/handles"
	"github.com/stephannykauane/projeto_it/backend/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/signup", handles.SignUp)
	mux.HandleFunc("/login", handles.Login)
	mux.Handle("/analise", middleware.Auth(http.HandlerFunc(handles.Analise)))
	mux.HandleFunc("/excel", handles.Excel)
	mux.Handle("/listar", middleware.Auth(http.HandlerFunc(handles.ListaCalculos)))
	mux.HandleFunc("/logout", handles.Logout)
	mux.Handle("/profile", middleware.Auth(http.HandlerFunc(handles.PerfilUsuario)))
}
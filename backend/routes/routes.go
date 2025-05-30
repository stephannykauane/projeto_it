package routes

import (
	"net/http"

	"github.com/stephannykauane/projeto_it/backend/handles"
	"github.com/stephannykauane/projeto_it/backend/headers"
	"github.com/stephannykauane/projeto_it/backend/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/signup", headers.SetHeaders(http.HandlerFunc(handles.SignUp)))
	mux.Handle("/login", headers.SetHeaders(http.HandlerFunc(handles.Login)))
	mux.Handle("/analise", headers.SetHeaders(middleware.Auth(http.HandlerFunc(handles.Analise))))
	mux.Handle("/excel", headers.SetHeaders(http.HandlerFunc(handles.Excel)))
	mux.Handle("/listar", middleware.Auth(http.HandlerFunc(handles.ListaCalculos)))
	mux.Handle("/logout", headers.SetHeaders(http.HandlerFunc(handles.Logout)))
	mux.Handle("/profile", middleware.Auth(http.HandlerFunc(handles.PerfilUsuario)))
}
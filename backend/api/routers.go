package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter (){
	router := mux.NewRouter()
	router.Use(setHeaders)
	router.HandleFunc("/login/{id}", GetLogin).Methods("GET")
	router.HandleFunc("/signup", PostSignUp).Methods("POST")
	//router.HandleFunc("/profile", GetProfile).Methods("GET")
	//router.HandleFunc("/profile", GetProfile).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))

}
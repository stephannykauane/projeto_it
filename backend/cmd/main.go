package main

import (
	"fmt"
	"net/http"
	datab "github.com/stephannykauane/projeto_it/backend/db"
	api "github.com/stephannykauane/projeto_it/backend/api"
	
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Backend funcionando!")
}

func main() {
    api.Routes()
	datab.Database() 
	http.HandleFunc("/", handler)
	fmt.Println("Servidor Go iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	db "github.com/stephannykauane/projeto_it/backend/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Backend funcionando!")
}

func main() {
	db.FuncaoPulbica()
	http.HandleFunc("/", handler)
	fmt.Println("Servidor Go iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

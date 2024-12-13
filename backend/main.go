package main

import (
	"fmt"
	"net/http"

	datab "github.com/stephannykauane/projeto_it/backend/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Backend funcionando!")
}

func main() {
    datab.database()
	http.HandleFunc("/", handler)
	fmt.Println("Servidor Go iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"github.com/stephannykauane/projeto_it/backend/db"
	"github.com/stephannykauane/projeto_it/backend/headers"
	"github.com/stephannykauane/projeto_it/backend/routes"
	"net/http"
)

func main() {

	db.Database()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend funcionando!")
	})

	fmt.Println("Servidor Go iniciado em http://localhost:8080")
	
	err := http.ListenAndServe(":8080", headers.SetHeaders(mux))
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}

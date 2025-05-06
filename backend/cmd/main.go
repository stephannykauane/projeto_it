package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/stephannykauane/projeto_it/backend/db"
	"github.com/stephannykauane/projeto_it/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}
	db.Database()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend funcionando!")
	})
	
	fmt.Println("Servidor Go iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

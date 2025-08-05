package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/stephannykauane/projeto_it/backend/db"
	"github.com/stephannykauane/projeto_it/backend/routes"
)

func main() {
	db.Database()
	db.RunMigrations()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Backend funcionando!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Servidor Go iniciado na porta " + port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
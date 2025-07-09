package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
)

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func init() {
	
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Aviso: erro ao carregar o arquivo .env. funcionando com os default values.")
		}
	}

	host = getEnv("DB_HOST", "localhost")
	user = getEnv("DB_USER", "calagem")
	password = getEnv("DB_PASSWORD", "admin")
	dbname = getEnv("DB_NAME", "calagem")

	portStr := getEnv("DB_PORT", "5432")
	var err error
	port, err = strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Erro ao converter DB_PORT: %v", err)
	}
}


func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func Database() {
	db := ConnectDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar com o banco de dados.")
		panic(err)
	}
	fmt.Println("Conexão com o banco de dados realizada com sucesso!")
}

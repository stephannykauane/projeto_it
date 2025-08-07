package db

import (
	"database/sql"
	"fmt"
	"log"
	// "net/url"
	"os"
	"strconv"

	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

// func RunMigrations() {

// 	escapePassword := url.QueryEscape(password)
// 	escapedUser := url.QueryEscape(user)
// 	escapedHost := url.QueryEscape(host)
// 	escapedDBname := url.QueryEscape(dbname)

// 	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
// 		escapedUser, escapePassword, escapedHost, port, escapedDBname)

// 	m, err := migrate.New(
// 		"file://internal/database/migrations", 
// 		dsn,
// 	)
// 	if err != nil {
// 		log.Fatalf("Erro ao iniciar migrator: %v", err)
// 	}

// 	if err := m.Up(); err != nil && err.Error() != "no change" {
// 		log.Fatalf("Erro ao aplicar migrations: %v", err)
// 	}

// 	log.Println("✅ Migrations aplicadas com sucesso.")
// }

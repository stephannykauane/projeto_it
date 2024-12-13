package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" 
)

const (
	host     = "localhost"
	port     = 5432
	user     = "calagem"
	password = "admin"
	dbname   = "calagem"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	return db
}

func Database() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot ping the database")
		panic(err)
	}
	db.Close()
	fmt.Println("Successfully connected with database!")
}

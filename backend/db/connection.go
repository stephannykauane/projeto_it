package db
import (
	"fmt"
	"database/sql"
)


const (
	host = "localhost"
	port = 5432
	user = "calagem"
	password = "12345"
	dbname = "calagem"
)

func connectdb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	return db
}
func database() {

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
	fmt.Println("Successfully connected!")

}





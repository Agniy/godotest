package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func mainConroler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi there,i wrote to test table number - 1")
}

func addTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
		`numbers("id" SERIAL PRIMARY KEY,` +
		`"number" varchar(50))`)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success create table numbers")
	}
}

func addNumber(db *sql.DB) {
	_, err := db.Exec("INSERT INTO numbers VALUES (default,$1)",
		"1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success insert to table numbers - 1")
	}
}

func main() {

	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	pgHost := os.Getenv("PGHOST")
	dbName := os.Getenv("PGDATABASE")

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		dbUser, dbPassword, pgHost, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	addTable(db)
	addNumber(db)

	http.HandleFunc("/", mainConroler)
	http.ListenAndServe(":8080", nil)
}

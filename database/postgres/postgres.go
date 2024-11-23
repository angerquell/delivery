package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil{
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

  }
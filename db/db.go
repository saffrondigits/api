package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "127.0.0.1"
	PORT     = "3333"
	USER     = "root"
	PASS     = "Secret"
	DATABASE = "golang_class"
)

func ConnectTOTheDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", USER, PASS, HOST, PORT, DATABASE)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

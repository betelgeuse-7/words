package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	PG_CONN_STR := os.Getenv("PG_CONN_STR")

	db, err = sql.Open("postgres", PG_CONN_STR)
	if err != nil {
		panic(err)
	}

	return db.Ping()
}

package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=go_db sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Connected to Database")
}

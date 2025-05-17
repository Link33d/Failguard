package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase() {

	connection_url := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", connection_url)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	log.Println("✅ Connected to PostgreSQL")

}

package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	//import for psql connection
	_ "github.com/lib/pq"
)

var con *sql.DB

// Connect Database
func Connect() *sql.DB {
	// Connect to the DB, panic if failed
	psqlURL := getPsqlURL("PSQL")
	db, err := sql.Open("postgres", psqlURL)
	if err != nil {
		fmt.Println(`Could not connect to db`)
	}
	log.Println("Connected to the database")
	con = db
	return db
}

func getPsqlURL(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

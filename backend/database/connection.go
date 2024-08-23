package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB is a global variable to hold the database connection pool
var DB *sqlx.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	// Database connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Open the database connection
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// Ping the database to check if the connection is successful
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}

	log.Println("Successfully connected to the database!")
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database: %q", err)
	}
}

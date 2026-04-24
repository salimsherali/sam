package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DB is a global variable to hold database connection
var DB *sql.DB

// Connect initializes and verifies database connection
func Connect() {

	// Build DSN (Data Source Name) from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open connection (does not establish immediately)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	// Verify connection with Ping
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	// Assign to global DB variable
	DB = db

	log.Println("Database connected successfully")
}
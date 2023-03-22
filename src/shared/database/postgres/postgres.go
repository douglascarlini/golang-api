package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB) {

	// Connection settings
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	// Crete the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

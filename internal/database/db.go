package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Default connection string, can be overridden by env var DATABASE_URL
	connStr := "postgres://postgres:postgres@localhost:5432/santaverse?sslmode=disable"
	if url := os.Getenv("DATABASE_URL"); url != "" {
		connStr = url
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database successfully")
}

func InitSchema(schemaPath string) {
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

	log.Println("Database schema initialized and seeded")
}

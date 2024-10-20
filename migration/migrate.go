package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUserName := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE") // e.g., "disable" or "require"
	// Build PostgreSQL connection string
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUserName, dbPass, dbHost, dbPort, dbName, dbSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Path to migration files
		"postgres", driver,
	)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Migrations applied successfully!")
}

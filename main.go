package main

import (
	"database/sql"
	"fmt"
	"inventory/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	db, err := loadDb()
	if err != nil {
		log.Fatal("Error loading database:", err)
		return
	}
	defer db.Close()

	r := mux.NewRouter()

	// protected := r.PathPrefix("/").Subrouter()
	// protected.Use(authMidd)
	r.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(db, w, r)
	}).Methods("POST")
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

}

func loadDb() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUserName := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE") // e.g., "disable" or "require"
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUserName, dbPass, dbHost, dbPort, dbName, dbSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}

// main.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"user-service/handler"
	"user-service/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create users table if not exists
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			updated_at INTEGER NOT NULL
		)
	`); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepositoryUserService(db)
	restHandler := handler.NewRestHandler(repo)

	// Set up routes
	http.HandleFunc("/users", restHandler.CreateOrGetAllUser)
	http.HandleFunc("/users/", restHandler.GetUserByID)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

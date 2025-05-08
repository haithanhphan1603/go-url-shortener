package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"url-shortener/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		 original_url TEXT NOT NULL,
		code TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		panic(err)
	}
	fmt.Println("✅ SQLite works!")
	return db
}

func main() {
	db := initDatabase()
	h := &handlers.Handler{DB: db}

	http.HandleFunc("/shorten", h.ShortenURLHandler)
	http.HandleFunc("/", h.RedirectHandler)

	fmt.Println("🚀 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

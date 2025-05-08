package db

import (
	"database/sql"
	"url-shortener/models"
)

func FetchURL(code string, db *sql.DB) (models.URL, error) {
	var url models.URL
	query := `SELECT id, original_url,code, created_at FROM urls WHERE code = ?`
	err := db.QueryRow(query, code).Scan(&url.ID, &url.OriginalURL, &url.Code, &url.CreatedAt)
	if err != nil {
		return models.URL{}, err
	}
	return url, nil
}

func InsertURL(url models.URL, db *sql.DB) (int64, error) {
	query := `INSERT INTO urls (original_url, code) VALUES (?, ?)`
	result, err := db.Exec(query, url.OriginalURL, url.Code)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

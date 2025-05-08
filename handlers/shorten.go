package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"url-shortener/db"
	"url-shortener/models"
	"url-shortener/utils"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var urlRequest struct {
		Url string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&urlRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = utils.ValidateUrl(urlRequest.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		http.Error(w, "Error generating code", http.StatusInternalServerError)
		return
	}

	newUrl := models.URL{
		OriginalURL: urlRequest.Url,
		Code:        code,
	}

	_, err = db.InsertURL(newUrl, h.DB)
	if err != nil {
		http.Error(w, "Error saving URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"short_url": "http://localhost:8080/" + code,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/"):]

	url, err := db.FetchURL(code, h.DB)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

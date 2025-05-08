package models

type URL struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	Code        string `json:"code"`
	CreatedAt   string `json:"created_at"`
}

package domain

import "time"

type Link struct {
	ID          int
	OriginalURL string
	ShortCode   string
	Clicks      int
	CreatedAt   time.Time
	ExpiresAt   time.Time
	IsActive    int
}

type LinkResponse struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	IsActive    int       `json:"is_active"`
}

type CreateLinkRequest struct {
	Original_URL string `json:"original_url"`
}

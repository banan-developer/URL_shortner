package domain

type Link struct {
	ID          int
	OriginalURL string
	ShortCode   string
	Clicks      int
	CreatedAt   string
	ExpiresAt   string
	IsActive    int
	UserID      int
}

type LinkResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	Clicks      int    `json:"clicks"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
	IsActive    int    `json:"is_active"`
}

type CreateLinkRequest struct {
	ID           int    `json:"id"`
	Original_URL string `json:"original_url"`
	IsAcitve     int    `json:"is_active"`
}

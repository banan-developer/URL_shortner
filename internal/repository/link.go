package repository

import (
	"URL_shortner/internal/domain"
	"database/sql"
)

type LinksRepo struct {
	db *sql.DB
}

func NewLinksRepo(db *sql.DB) *LinksRepo {
	return &LinksRepo{
		db: db,
	}
}

func (r *LinksRepo) GetLinkByID(UserID int) (*domain.LinkResponse, error) {
	var Links domain.LinkResponse
	err := r.db.QueryRow("SELECT original_url, short_url, clicks, created_at, expires_at, is_active FROM links WHERE users_id = ?", UserID).Scan(&Links.OriginalURL, &Links.ShortURL, &Links.Clicks, &Links.CreatedAt, &Links.ExpiresAt, &Links.IsActive)
	if err != nil {
		return nil, err
	}
	return &Links, nil
}

func (r *LinksRepo) CreateLink(link *domain.Link) error {
	var userID interface{}

	if link.UserID > 0 {
		userID = link.UserID
	} else {
		userID = nil
	}
	_, err := r.db.Exec("INSERT INTO links (original_url, short_url, clicks, expires_at, is_active, users_id) VALUES (?, ?, ?, ?, ?, ?)", link.OriginalURL, link.ShortCode, link.Clicks, link.ExpiresAt, link.IsActive, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *LinksRepo) GetLinkByShortlink(shortCode string) (*domain.CreateLinkRequest, error) {
	var Link domain.CreateLinkRequest
	err := r.db.QueryRow("SELECT original_url FROM links WHERE short_url = ?", shortCode).Scan(&Link.Original_URL)
	if err != nil {
		return nil, err
	}
	return &Link, nil
}

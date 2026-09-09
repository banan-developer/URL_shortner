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

// func (r *LinksRepo) GetLinkByID(UserID int) (*domain.LinkResponse, error) {
// 	var Links domain.LinkResponse
// 	err := r.db.QueryRow("SELECT original_url, short_url, clicks, created_at, expires_at, is_active FROM links WHERE users_id = ?", UserID).Scan(&Links.OriginalURL, &Links.ShortURL, &Links.Clicks, &Links.CreatedAt, &Links.ExpiresAt, &Links.IsActive)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Links, nil
// }

func (r *LinksRepo) GetLinksByUserID(UserID int) ([]domain.LinkResponse, error) {
	rows, err := r.db.Query("SELECT id, original_url, short_url, clicks, created_at, expires_at, is_active FROM links WHERE users_id = ?", UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var links []domain.LinkResponse

	for rows.Next() {
		var link domain.LinkResponse
		if err := rows.Scan(&link.ID, &link.OriginalURL, &link.ShortURL, &link.Clicks, &link.CreatedAt, &link.ExpiresAt, &link.IsActive); err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	if links == nil {
		links = []domain.LinkResponse{}
	}
	return links, err
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
	err := r.db.QueryRow("SELECT id, original_url, is_active FROM links WHERE short_url = ?", shortCode).Scan(&Link.ID, &Link.Original_URL, &Link.IsAcitve)
	if err != nil {
		return nil, err
	}
	return &Link, nil
}

func (r *LinksRepo) DeleteLinkByID(LinkID int, UserID int) error {
	_, err := r.db.Exec("DELETE FROM links WHERE id = ? AND users_id = ?", LinkID, UserID)
	return err
}

func (r *LinksRepo) UpdateIsActive(LinkID int, Action int) error {
	_, err := r.db.Exec("UPDATE links SET is_active = ? WHERE id = ?", Action, LinkID)
	return err
}

func (r *LinksRepo) IncrementClicks(LinkID int) error {
	_, err := r.db.Exec("UPDATE links SET clicks = clicks + 1 WHERE id = ?", LinkID)
	return err
}

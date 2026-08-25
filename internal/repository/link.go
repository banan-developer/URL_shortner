package repository

import (
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

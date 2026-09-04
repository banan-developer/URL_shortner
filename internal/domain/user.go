package domain

import "time"

type User struct {
	ID       int
	UserName string
	Email    string
	Plan     string
	Password string
}

type UserResponse struct {
	UserName    string    `json:"username"`
	Plan        string    `json:"plan"`
	Created_at  time.Time `json:"created_at"`
	AllLinks    int       `json:"all_links"`
	ActiveLinks int       `json:"active_links"`
}

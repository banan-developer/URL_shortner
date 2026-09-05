package repository

import (
	"URL_shortner/internal/domain"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(User *domain.User, hashPassword string) error {
	_, err := u.db.Exec("INSERT INTO users (username, email, plan, password) VALUES (?, ?, ?, ?)", User.UserName, User.Email, User.Plan, hashPassword)
	return err
}

func (u *UserRepo) GetUserByLogin(Login string) (int, string, error) {
	var UserID int
	var PasswordFromDB string

	err := u.db.QueryRow("SELECT id, password FROM users WHERE email = ?", Login).Scan(&UserID, &PasswordFromDB)
	if err != nil {
		return 0, "", err
	}
	return UserID, PasswordFromDB, nil
}

func (u *UserRepo) GetUserByID(User *domain.UserResponse, UserID int) error {
	err := u.db.QueryRow("SELECT id, username, plan, created_at FROM users WHERE id = ?", UserID).Scan(&User.UserID, &User.UserName, &User.Plan, User.Created_at)
	return err
}

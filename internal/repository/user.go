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
	err := u.db.QueryRow(`
		SELECT 
			u.id,
			u.username,
			u.plan,
			u.created_at,
			COUNT(l.id) AS all_links,
			SUM(CASE WHEN l.is_active = 1 THEN 1 ELSE 0 END) AS active_links
		FROM users u
		LEFT JOIN links l ON u.id = l.users_id
		WHERE u.id = ?
		GROUP BY u.id, u.username, u.plan, u.created_at
	`, UserID).Scan(
		&User.UserID,
		&User.UserName,
		&User.Plan,
		&User.Created_at,
		&User.AllLinks,
		&User.ActiveLinks,
	)
	return err
}

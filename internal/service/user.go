package service

import (
	"URL_shortner/internal/domain"
	"URL_shortner/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (s *UserService) RegistratingUser(User *domain.User) error {
	var hashedPassword, err = HashPassword(User.Password)
	if err != nil {
		return err
	}
	User.Plan = "standart"
	return s.repo.CreateUser(User, hashedPassword)
}

func (s *UserService) LoginUser(Login string, password string) (int, error) {
	if Login == "" {
		return 0, errors.New("ошибка при получении почты пользователя")
	}
	UserID, PasswordFromdb, err := s.repo.GetUserByLogin(Login)
	if err != nil {
		return 0, errors.New("Invalid credentials")
	}

	HashError := bcrypt.CompareHashAndPassword(
		[]byte(PasswordFromdb),
		[]byte(password),
	)
	if HashError != nil {
		return 0, errors.New("Ошибка авторизации")
	}
	return UserID, nil
}

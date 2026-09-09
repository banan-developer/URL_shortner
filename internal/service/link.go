package service

import (
	"URL_shortner/internal/domain"
	"URL_shortner/internal/generation"
	"URL_shortner/internal/repository"
	"errors"
	"time"
)

type LinksService struct {
	repo *repository.LinksRepo
}

func NewLinksService(repo *repository.LinksRepo) *LinksService {
	return &LinksService{
		repo: repo,
	}
}

func (s *LinksService) GetLinksByUserID(UserID int) ([]domain.LinkResponse, error) {
	if UserID < 0 {
		return nil, errors.New("invalid UserID")
	}
	return s.repo.GetLinksByUserID(UserID)
}

func (s *LinksService) CreateLink(OriginalURL string, UserID int) (string, error) {
	var shortURL, err = generation.GenerationCode(6)
	if err != nil {
		return "", err
	}
	link := &domain.Link{
		OriginalURL: OriginalURL,
		ShortCode:   shortURL,
		Clicks:      0,
		ExpiresAt:   time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05"),
		IsActive:    1,
	}

	if UserID > 0 {
		link.UserID = UserID
	}

	err = s.repo.CreateLink(link)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *LinksService) GetLinkByShortlink(ShortCode string) (*domain.CreateLinkRequest, error) {
	if ShortCode == "" {
		return nil, errors.New("Ошибка при получении короткой ссылки")
	}

	return s.repo.GetLinkByShortlink(ShortCode)
}

func (s *LinksService) DeleteLinkByID(LinkID int, UserID int) error {
	if LinkID <= 0 {
		return errors.New("неправильный id ссылки")
	}
	return s.repo.DeleteLinkByID(LinkID, UserID)
}

func (s *LinksService) UpdateIsActive(LinkID, Active int) error {
	if LinkID <= 0 {
		return errors.New("неправильный id ссылки")
	}
	return s.repo.UpdateIsActive(LinkID, Active)
}

func (s *LinksService) IncrementClicks(LinkID int) error {
	return s.repo.IncrementClicks(LinkID)
}

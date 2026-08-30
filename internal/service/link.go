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

func (s *LinksService) GetLinkByID(UserID int) (*domain.LinkResponse, error) {
	if UserID < 0 {
		return nil, errors.New("invalid UserID")
	}
	return s.repo.GetLinkByID(UserID)
}

func (s *LinksService) CreateLink(OriginalURL string) (string, error) {
	var shortURL, err = generation.GenerationCode(6)
	if err != nil {
		return "", err
	}
	link := &domain.Link{
		OriginalURL: OriginalURL,
		ShortCode:   shortURL,
		Clicks:      0,
		ExpiresAt:   time.Now().AddDate(0, 0, 7),
		IsActive:    1,
	}
	err = s.repo.CreateLink(link)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

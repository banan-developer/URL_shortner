package service

import "URL_shortner/internal/repository"

type LinksService struct {
	repo *repository.LinksRepo
}

func NewLinksService(repo *repository.LinksRepo) *LinksService {
	return &LinksService{
		repo: repo,
	}
}

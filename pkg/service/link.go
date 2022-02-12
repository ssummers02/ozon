package service

import (
	"ozon/pkg/repository"
	"ozon/pkg/restmodel"
	"time"
)

type LinkService struct {
	repo repository.Link
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) Create(link restmodel.Link) (restmodel.Link, error) {
	shortLink, err := s.repo.GetByLink(link)
	if err != nil {
		link.CreatedAt = time.Now().UTC()
		link.ShortLink = generShortLink()
		return s.repo.Create(link)
	}

	return shortLink, nil
}

func (s *LinkService) GetByShortLink(link restmodel.Link) (restmodel.Link, error) {
	return s.repo.GetByShortLink(link)
}

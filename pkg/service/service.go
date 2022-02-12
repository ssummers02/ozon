package service

import (
	"ozon/pkg/repository"
	"ozon/pkg/restmodel"
)

type Link interface {
	Create(link restmodel.Link) (restmodel.Link, error)
	GetByShortLink(link restmodel.Link) (restmodel.Link, error)
}

type Service struct {
	Link Link
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Link: NewLinkService(repos.Link),
	}
}

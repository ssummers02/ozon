package repository

import (
	"ozon/pkg/restmodel"

	"github.com/gocraft/dbr"
)

type Link interface {
	Create(link restmodel.Link) (restmodel.Link, error)
	GetByShortLink(link restmodel.Link) (restmodel.Link, error)
	GetByLink(link restmodel.Link) (restmodel.Link, error)
}

type Repository struct {
	Link
}

func NewRepository(db *dbr.Connection) *Repository {
	return &Repository{
		Link: NewLinkPostgres(db),
	}
}

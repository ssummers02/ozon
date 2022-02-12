package repository

import (
	"ozon/pkg/restmodel"

	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type LinkPostgres struct {
	db *dbr.Connection
}

func NewLinkPostgres(db *dbr.Connection) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (l *LinkPostgres) Create(link restmodel.Link) (restmodel.Link, error) {
	err := l.db.NewSession(nil).InsertInto(linkTable).
		Returning("id").
		Pair("link", link.Link).
		Pair("short_link", link.ShortLink).
		Pair("created_at", link.CreatedAt).
		Load(&link.ID)
	if err != nil {
		return restmodel.Link{}, errors.WithStack(err)
	}

	return link, nil
}

func (l *LinkPostgres) GetByShortLink(link restmodel.Link) (restmodel.Link, error) {
	err := l.db.NewSession(nil).
		Select("link", "short_link", "created_at", "id").
		From(linkTable).
		Where("short_link = ?", link.ShortLink).
		LoadOne(&link)
	if err != nil {
		return restmodel.Link{}, errors.WithStack(err)
	}

	return link, nil
}

func (l *LinkPostgres) GetByLink(link restmodel.Link) (restmodel.Link, error) {
	err := l.db.NewSession(nil).
		Select("link", "short_link", "created_at", "id").
		From(linkTable).
		Where("link = ?", link.Link).
		LoadOne(&link)
	if err != nil {
		return restmodel.Link{}, errors.WithStack(err)
	}

	return link, nil
}

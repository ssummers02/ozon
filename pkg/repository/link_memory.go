package repository

import (
	"ozon/pkg/restmodel"
	"sync"

	"github.com/pkg/errors"
)

type LinkMemory struct {
	db map[string]string
	mu sync.RWMutex
}

func (l *LinkMemory) Create(link restmodel.Link) (restmodel.Link, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.db[link.Link] = link.ShortLink
	return link, nil
}

func (l *LinkMemory) GetByShortLink(link restmodel.Link) (restmodel.Link, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for key, value := range l.db {
		if value == link.ShortLink {
			link.Link = key
			return link, nil
		}
	}
	return restmodel.Link{}, errors.New("not found")
}

func (l *LinkMemory) GetByLink(link restmodel.Link) (restmodel.Link, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	i, ok := l.db[link.Link]
	if !ok {
		return restmodel.Link{}, errors.New("not found")
	}

	link.ShortLink = i

	return link, nil
}

func NewLinkMemory() *LinkMemory {
	return &LinkMemory{
		db: make(map[string]string),
		mu: sync.RWMutex{},
	}
}

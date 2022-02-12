package repository

import (
	"github.com/hashicorp/go-memdb"
)

type LinkMemory struct {
	db *memdb.DBSchema
}

func NewLinkMemory(db *memdb.DBSchema) *LinkMemory {
	return &LinkMemory{db: db}
}

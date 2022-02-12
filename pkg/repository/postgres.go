package repository

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	DBSchema string
}

const maxOpenConns = 10

const (
	linkTable = "link"
)

func NewPostgresDB(cfg Config) (*dbr.Connection, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	log.Println(dsn)
	p, err := dbr.Open(cfg.DBSchema, dsn, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// verifying that it's alive at the beginning
	err = p.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// a maximum of 10 concurrent connections; might be changed later if needed
	p.SetMaxOpenConns(maxOpenConns)

	return p, nil
}

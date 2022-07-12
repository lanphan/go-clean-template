// Package postgres implements postgres connection.
package postgres

import (
	"github.com/ironsail/whydah-go-clean-template/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

// Postgres -.
type Postgres db.Session

// New -.
func New(cfg *config.Config) (Postgres, error) {
	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: cfg.PG.DbName,
		Host:     cfg.PG.Host,
		User:     cfg.PG.User,
		Password: cfg.PG.Password,
	})

	if err != nil {
		return nil, err
	}

	newVar := sess.(Postgres)
	return newVar, nil
}

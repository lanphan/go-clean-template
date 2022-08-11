// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"

	"github.com/ironsail/whydah-go-clean-template/pkg/postgres"
	"github.com/upper/db/v4"
)

// UserStore represents a pool of users
type UserStore struct {
	db.Collection
}

// Users initializes a UserStore
func Users(sess postgres.Postgres) *UserStore {
	return &UserStore{sess.Collection("users")}
}

// User
type User struct {
	ID            string    `db:"id,omitempty"`
	WalletAddress string    `db:"wallet_address"`
	Reward        uint      `db:"reward,omitempty"`
	CreatedAt     time.Time `db:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at,omitempty"`
}

func (user *User) Store(sess db.Session) db.Store {
	return sess.Collection("users")
}

func (user *User) ToRecord() db.Record {
	return db.Record(user)
}

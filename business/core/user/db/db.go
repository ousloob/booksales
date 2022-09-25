// Package db contains user related CRUD functionality.
package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/oussamm/bookstore/vendor/github.com/jmoiron/sqlx"
)

var ErrUniqueEmail = errors.New("email is not unique")

// Store manages the set of APIs for user access.
type Store struct {
	log *log.Logger
	db  *sqlx.DB
}

// NewStore constructs a data for api access.
func NewStore(log *log.Logger, db *sqlx.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new user into the database.
func (s Store) Create(ctx context.Context, usr User) error {
	const q = `
	INSERT into users (
        user_id,
        name,
        email,
        password_hash,
        roles,
        date_created,
        date_updated
    )
	VALUES (
        :user_id,
        :name,
        :email,
        :password_hash,
        :roles,
        :date_created,
        date_updated
    )`

	if _, err := sqlx.NamedExecContext(ctx, s.db, q, usr); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}

// Package userdb contains user related CRUD functionality.
package userdb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var ErrUniqueEmail = errors.New("email is not unique")

// Bus manages the set of APIs for user access.
type Bus struct {
	log *log.Logger
	db  *sqlx.DB
}

// NewBus constructs a data for api access.
func NewBus(log *log.Logger, db *sqlx.DB) *Bus {
	return &Bus{
		log: log,
		db:  db,
	}
}

// Create inserts a new user into the database.
func (b *Bus) Create(ctx context.Context, usr User) error {
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

	if _, err := sqlx.NamedExecContext(ctx, b.db, q, usr); err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}

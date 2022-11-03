// Package user provides a functionality of a core business API for user.
package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/oussamm/bookstore/business/core/user/db"
	"github.com/oussamm/bookstore/business/sys/database"
	"github.com/oussamm/bookstore/business/sys/validate"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Core manages the set of APIs for user access.
type Core struct {
	store db.Store
}

func NewCore(log *log.Logger, sqlxDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, sqlxDB),
	}
}

// Create inserts a new user into the database.
func (c Core) Create(ctx context.Context, nu NewUser, now time.Time) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash: %w", err)
	}

	dbUser := db.User{
		ID:           validate.GenerateID(),
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: hash,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := c.store.Create(ctx, dbUser); err != nil {
		if errors.Is(err, database.ErrDBDuplicatedEntry) {
			return User{}, fmt.Errorf("create: %w", db.ErrUniqueEmail)
		}
	}

	return User{
		ID:           dbUser.ID,
		Name:         dbUser.Name,
		Email:        dbUser.Email,
		Roles:        dbUser.Roles,
		PasswordHash: dbUser.PasswordHash,
		DateCreated:  dbUser.DateCreated,
		DateUpdated:  dbUser.DateUpdated,
	}, nil
}

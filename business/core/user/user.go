// Package user provides a functionality of a core business API for user.
package user

import "github.com/oussamm/bookstore/business/core/user/db"

// Core manages the set of APIs for user access.
type Core struct {
	store db.Store
}

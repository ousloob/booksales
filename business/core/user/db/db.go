// Package db contains user related CRUD functionality.
package db

import "log"

// Store manages the set of APIs for user access.
type Store struct {
	log *log.Logger
}

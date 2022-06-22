package user

import "time"

// User represent the structure we need for moving data
// between the app and the database.
type User struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	ID           string    `json:"user_id"`
	Roles        []string  `json:"roles"`
	PasswordHash []byte    `json:"password_hash"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
}

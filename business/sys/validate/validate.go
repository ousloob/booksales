// Package validate contains the support for validating models.
package validate

import "github.com/google/uuid"

// GenerateID generate a unique ID for entities.
func GenerateID() string {
	return uuid.NewString()
}

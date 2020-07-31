package utils

import (
	"github.com/google/uuid"
)

// New return new unique ID
func NewID() string {
	return uuid.New().String()
}

// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"github.com/google/uuid"
)

type Users struct {
	ID       uuid.UUID
	Name     string
	Lastname string
	Username string
	Country  string
}

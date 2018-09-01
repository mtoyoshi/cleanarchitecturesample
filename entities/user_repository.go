package entities

import (
	"github.com/satori/go.uuid"
)

type UserRepository interface {
	Store(user User) error
	Find(id uuid.UUID) (User, error)
}

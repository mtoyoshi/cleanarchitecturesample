package entities

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type User struct {
	Id   uuid.UUID
	Name string
	Age  int
}

func NewUser(name string, age int) (User, error) {
	var errMsgs []string

	if len(name) <= 0 || len(name) > 100 {
		errMsgs = append(errMsgs, "name length must be between 1-100")
	}
	if age < 18 || age > 80 {
		errMsgs = append(errMsgs, "age range must be between 1-100")
	}
	if len(errMsgs) > 0 {
		return User{}, errors.New(strings.Join(errMsgs, ", "))
	}

	return User{
		Id:   uuid.NewV4(),
		Name: name,
		Age:  age,
	}, nil
}

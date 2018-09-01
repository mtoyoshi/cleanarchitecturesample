package entities

import (
	"fmt"
	"github.com/mtoyoshi/cleanarchitecturesample/entities"
	"github.com/satori/go.uuid"
)

type UserRepositoryImpl struct{}

func (repo UserRepositoryImpl) Store(user entities.User) error {
	fmt.Printf("mock: stored user %s(%d)", user.Name, user.Age)
	return nil
}

func (repo UserRepositoryImpl) Find(id uuid.UUID) (entities.User, error) {
	return entities.User{
		Id:   id,
		Name: "mock: this is mock-user",
		Age:  20,
	}, nil
}

var UserRepositoryImplInstance = UserRepositoryImpl{}

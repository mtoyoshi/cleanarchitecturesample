package usecases

import (
	"github.com/mtoyoshi/cleanarchitecturesample/entities"
)

type RegisterUserUseCase struct {
	UserRepository entities.UserRepository
}

func (uc RegisterUserUseCase) DoRun(param interface{}) UseCaseResult {
	p := param.(RegisterUserParam)
	user, err := entities.NewUser(p.Name, p.Age)
	if err != nil {
		return uc.makeResult(user, err)
	}

	if err := uc.UserRepository.Store(user); err != nil {
		return uc.makeResult(user, err)
	}

	return uc.makeResult(user, err)
}

func (uc RegisterUserUseCase) makeResult(user entities.User, e error) UseCaseResult {
	return UseCaseResult{
		Value: user,
		Error: e,
	}
}

type RegisterUserParam struct {
	Name string
	Age  int
}

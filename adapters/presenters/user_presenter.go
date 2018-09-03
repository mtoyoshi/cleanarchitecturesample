package presenters

import (
	"github.com/mtoyoshi/cleanarchitecturesample/entities"
)

type UserPresenterForRegistering struct{}

func (p UserPresenterForRegistering) DoRun(result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": result.(entities.User).Id,
	}
}

var UserPresenterInstance = UserPresenterForRegistering{}

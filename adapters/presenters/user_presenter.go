package presenters

import (
	"github.com/mtoyoshi/cleanarchitecturesample/entities"
)

type UserPresenter struct{}

func (p UserPresenter) DoRun(result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": result.(entities.User).Id,
	}
}

var UserPresenterInstance = UserPresenter{}

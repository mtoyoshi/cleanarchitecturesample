package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mtoyoshi/cleanarchitecturesample/adapters/entities"
	"github.com/mtoyoshi/cleanarchitecturesample/adapters/presenters"
	"github.com/mtoyoshi/cleanarchitecturesample/usecases"
	"strconv"
)

func RegisterUser(c *gin.Context) {
	inputPort := getRegisterUserInputPort()
	outputPort := getUserJsonPresenter(c)

	c.Request.ParseForm()
	name := c.Request.Form["name"][0]
	ageStr := c.Request.Form["age"][0]
	age, _ := strconv.Atoi(ageStr)

	param := usecases.RegisterUserParam{
		Name: name,
		Age:  age,
	}

	inputPort.Run(param, outputPort)
}

var getRegisterUserInputPort func() usecases.InputPort = (func() func() usecases.InputPort {
	inputPort := usecases.InputPort{
		UseCase: usecases.RegisterUserUseCase{
			UserRepository: entities.UserRepositoryImplInstance,
		},
	}
	return func() usecases.InputPort {
		return inputPort
	}
})()

func getUserJsonPresenter(c *gin.Context) presenters.JsonPresenter {
	return presenters.JsonPresenter{
		Ctx:            c,
		InnerPresenter: presenters.UserPresenterInstance,
	}
}

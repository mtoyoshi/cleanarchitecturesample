package presenters

import (
	"github.com/gin-gonic/gin"
	"github.com/mtoyoshi/cleanarchitecturesample/usecases"
)

type InnerPresenter interface {
	DoRun(result interface{}) map[string]interface{}
}

type JsonPresenter struct {
	Ctx            *gin.Context
	InnerPresenter InnerPresenter
}

func (p JsonPresenter) Run(result usecases.UseCaseResult) {
	if result.HasError() {
		p.Ctx.JSON(200, gin.H{"erorr": result.Error.Error()})
	} else {
		p.Ctx.JSON(200, gin.H(p.InnerPresenter.DoRun(result.Value)))
	}
}

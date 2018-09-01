package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mtoyoshi/cleanarchitecturesample/adapters/controllers"
)

func main() {
	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	r.Run()
}

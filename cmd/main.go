package main

import (
	"github.com/gin-gonic/gin"
	"qyx/pkg/start"
)

func main() {
	app := CreateApp()
	Start(app)
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}

func Start(ginStart *start.GinStart) error {
	return ginStart.Gin.Run(":80")

}

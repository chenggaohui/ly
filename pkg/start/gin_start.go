package start

import "github.com/gin-gonic/gin"

type Controllers func(g *gin.Engine)

type GinStart struct {
	Gin *gin.Engine
}

func NewGinStart(g *gin.Engine, controllers Controllers) *GinStart {
	controllers(g)
	return &GinStart{
		Gin: g,
	}
}

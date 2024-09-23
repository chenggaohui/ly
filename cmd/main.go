package main

import (
	"github.com/gin-gonic/gin"
	"ly/config"
	"ly/pkg/start"
)

func init() {
	config.InitConfig()
}

func main() {
	app := CreateApp()
	Start(app)
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	// 添加跨域中间件
	engine.Use(Cors())
	return engine
}

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}

func Start(ginStart *start.GinStart) error {
	return ginStart.Gin.Run(":8081")

}

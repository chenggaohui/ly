package controllers

import (
	"github.com/gin-gonic/gin"
	"qyx/pkg/start"
)

func NewController(
	helloController *CalculateController,
) start.Controllers {
	return func(r *gin.Engine) {
		r.GET("hello", helloController.Hello)
	}
}

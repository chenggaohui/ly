package controllers

import (
	"github.com/gin-gonic/gin"
	"qyx/internal/services"
)

type CalculateController struct {
	helloService services.HelloculateService
}

func (ctl CalculateController) Hello(c *gin.Context) {
	msg, err := ctl.helloService.Hello()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, msg)
	return
}

func NewCalculateController(
	helloService services.HelloculateService,
) *CalculateController {
	return &CalculateController{
		helloService: helloService,
	}
}

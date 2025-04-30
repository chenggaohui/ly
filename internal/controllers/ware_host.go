package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ly/internal/services"
	"ly/pkg/models"
	"strconv"
)

type WareHostController struct {
	WarehouseService services.WareHostService
}

func NewWareHostController(
	WarehouseService services.WareHostService,
) *WareHostController {
	return &WareHostController{
		WarehouseService: WarehouseService,
	}
}

func (w *WareHostController) GetWarehouseList(c *gin.Context) {
	ware := &models.WareHost{}
	err := c.BindQuery(ware)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	list, err := w.WarehouseService.GetAll(c.Request.Context(), ware.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
			"data":    list,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    list,
	})

}

func (w *WareHostController) CreateWarehouse(context *gin.Context) {
	ware := &models.WareHost{}
	err := context.BindJSON(ware)
	if err != nil {
		context.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = w.WarehouseService.Create(context.Request.Context(), ware)
	if err != nil {
		context.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
		})
		return
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func (w *WareHostController) UpdateWarehouse(context *gin.Context) {
	ware := &models.WareHost{}
	err := context.BindJSON(ware)
	if err != nil {
		context.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = w.WarehouseService.Update(context.Request.Context(), ware)
	if err != nil {
		context.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
		})
		return
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func (w *WareHostController) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	wareId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = w.WarehouseService.Delete(c.Request.Context(), wareId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

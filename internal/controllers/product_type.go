package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ly/internal/services"
	"ly/pkg/models"
	"strconv"
)

type ProductTypeController struct {
	productTypeService services.ProductTypeService
}

func NewProductTypeController(
	productTypeService services.ProductTypeService,
) *ProductTypeController {
	return &ProductTypeController{
		productTypeService: productTypeService,
	}
}

func (p *ProductTypeController) GetProductTypeList(c *gin.Context) {
	product := &models.ProductType{}
	err := c.BindQuery(product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	list, err := p.productTypeService.GetAll(c.Request.Context(), product.Name)
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

func (p *ProductTypeController) CreateProductType(c *gin.Context) {
	productType := &models.ProductType{}
	err := c.BindJSON(productType)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = p.productTypeService.Create(c.Request.Context(), productType)
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

func (p *ProductTypeController) UpdateProductType(c *gin.Context) {
	productType := &models.ProductType{}
	err := c.BindJSON(productType)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = p.productTypeService.Update(c.Request.Context(), productType)
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

func (p *ProductTypeController) DeleteProductType(c *gin.Context) {
	id := c.Param("id")
	productTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = p.productTypeService.Delete(c.Request.Context(), productTypeId)
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

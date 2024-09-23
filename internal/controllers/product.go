package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ly/internal/services"
	"ly/pkg/models"
	"ly/pkg/vo"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	product := &vo.ProductVo{}
	err := c.BindJSON(product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = p.productService.CreateProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
		})
		return
	}
}

func (p *ProductController) GetProductList(c *gin.Context) {
	product := &models.Product{}
	err := c.BindQuery(product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	pageNum := c.GetInt("pageNum")
	pageSize := c.GetInt("pageSize")
	list, total, err := p.productService.GetProductList(c.Request.Context(), product, pageNum, pageSize)
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
		"total":   total,
	})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	product := &vo.ProductVo{}
	err := c.BindJSON(product)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("参数错误),error:%s", err.Error()),
		})
		return
	}
	err = p.productService.UpdateProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("服务器错误,error:%s", err.Error()),
		})
		return
	}
}

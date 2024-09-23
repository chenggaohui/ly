package controllers

import (
	"github.com/gin-gonic/gin"
	"ly/pkg/start"
)

func NewController(
	productController *ProductController,
) start.Controllers {
	return func(r *gin.Engine) {
		g := r.Group("/api/v1")

		//product
		product := g.Group("/product")
		product.GET("/list", productController.GetProductList)
		product.POST("", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
	}
}

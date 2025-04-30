package controllers

import (
	"github.com/gin-gonic/gin"
	"ly/pkg/start"
)

func NewController(
	productController *ProductController,
	warehouseController *WareHostController,
	productTypeController *ProductTypeController,
) start.Controllers {
	return func(r *gin.Engine) {
		g := r.Group("/api/v1")

		//product
		product := g.Group("/product")
		product.GET("/list", productController.GetProductList)
		product.POST("", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)

		//warehouse
		warehouse := g.Group("/warehouse")
		warehouse.GET("/list", warehouseController.GetWarehouseList)
		warehouse.POST("", warehouseController.CreateWarehouse)
		warehouse.PUT("/:id", warehouseController.UpdateWarehouse)
		warehouse.DELETE("/:id", warehouseController.DeleteWarehouse)

		//product type
		productType := g.Group("/product_type")
		productType.GET("/list", productTypeController.GetProductTypeList)
		productType.POST("", productTypeController.CreateProductType)
		productType.PUT("/:id", productTypeController.UpdateProductType)
		productType.DELETE("/:id", productTypeController.DeleteProductType)
	}
}

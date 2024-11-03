package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controller *ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("", controller.CreateProduct)
		productGroup.GET("", controller.GetProducts)
		productGroup.GET("/:id", controller.GetProductById)
		productGroup.PUT("/:id", controller.UpdateProduct)
		productGroup.DELETE("/:id", controller.RemoveProduct)
	}
}

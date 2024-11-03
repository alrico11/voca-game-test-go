package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductModule(db *gorm.DB, router *gin.Engine) {
	repository := NewProductRepository(db)
	service := NewProductService(repository)
	controller := NewProductController(service)

	RegisterProductRoutes(router, controller)
}

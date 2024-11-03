package wallet

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupWalletModule(db *gorm.DB, router *gin.Engine) {
	repository := NewWalletRepository(db)
	service := NewWalletService(repository)
	controller := NewWalletController(service)

	RegisterWalletRoutes(router, controller)
}

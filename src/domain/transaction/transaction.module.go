// src/transaction/module.go
package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTransactionModule(db *gorm.DB, router *gin.Engine) {
	repo := NewTransactionRepository(db)
	service := NewTransactionService(repo)
	controller := NewTransactionController(service)

	RegisterTransactionRoutes(router, controller)
}

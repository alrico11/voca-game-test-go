package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthModule(db *gorm.DB, router *gin.Engine) {
	repository := NewAuthRepository(db)
	service := NewAuthService(repository)
	controller := NewAuthController(service)

	RegisterAuthRoutes(router, controller)
}

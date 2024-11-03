package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.Engine, controller *AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", controller.Register)
		authGroup.POST("/login", controller.Login)
	}
}

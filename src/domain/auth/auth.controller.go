package auth

import (
	"net/http"
	"tech-testing/src/models"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{service: service}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var createUserDTO CreateUserDTO
	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     createUserDTO.Name,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

	if err := ctrl.service.Register(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (ctrl *AuthController) Login(c *gin.Context) {

	var loginUserDTO LoginUserDTO
	if err := c.ShouldBindJSON(&loginUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.Login(loginUserDTO.Email, loginUserDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, AuthResponseDTO{Token: token})
}

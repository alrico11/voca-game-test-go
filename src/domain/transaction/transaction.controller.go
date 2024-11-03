package transaction

import (
	"net/http"
	"tech-testing/src/models"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *TransactionService
}

func NewTransactionController(service *TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var dto CreateTransactionDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := ctx.GetString("userID") // Ambil userID dari context, misalnya dari middleware auth
	wallet := &models.Wallet{
		ID:      "wallet123", // Ganti dengan wallet ID yang sebenarnya dari auth context atau database
		Balance: 5000,        // Misalnya ambil balance dari database
	}

	message, err := c.service.CreateTransaction(userID, &dto, wallet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

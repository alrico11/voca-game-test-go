// controller.go
package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	service *WalletService
}

func NewWalletController(service *WalletService) *WalletController {
	return &WalletController{service: service}
}

func (ctl *WalletController) CheckBalance(c *gin.Context) {
	userID := c.GetString("userID")
	wallet, err := ctl.service.CheckBalance(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wallet": wallet})
}

func (ctl *WalletController) Deposit(c *gin.Context) {
	userID := c.GetString("userID")
	dto, err := BindWalletDepositDTO(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wallet, err := ctl.service.Deposit(userID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wallet": wallet})
}

func (ctl *WalletController) Withdrawal(c *gin.Context) {
	userID := c.GetString("userID")
	dto, err := BindWalletWithdrawalDTO(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wallet, err := ctl.service.Withdrawal(userID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wallet": wallet})
}

func (ctl *WalletController) GetTransactionHistory(c *gin.Context) {
	userID := c.GetString("userID")
	dto, err := BindHistoryWalletQueryDTO(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactions, count, err := ctl.service.GetTransactionHistory(userID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transactions": transactions, "count": count})
}

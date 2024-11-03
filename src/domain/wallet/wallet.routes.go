package wallet

import (
	"tech-testing/src/domain/auth"

	"github.com/gin-gonic/gin"
)

func RegisterWalletRoutes(router *gin.Engine, walletController *WalletController) {
	walletGroup := router.Group("/wallet")
	{
		walletGroup.Use(auth.AuthMiddleware())
		walletGroup.POST("/deposit", walletController.Deposit)
		walletGroup.POST("/withdrawal", walletController.Withdrawal)
		walletGroup.GET("/balance", walletController.CheckBalance)
		walletGroup.GET("/transactions", walletController.GetTransactionHistory)
	}
}

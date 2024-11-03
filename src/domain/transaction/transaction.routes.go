// src/transaction/routes.go
package transaction

import "github.com/gin-gonic/gin"

func RegisterTransactionRoutes(router *gin.Engine, controller *TransactionController) {
	transactions := router.Group("/transactions")
	{
		transactions.POST("/", controller.CreateTransaction)
	}
}

// src/transaction/dto.go
package transaction

type CreateTransactionDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Qty       int    `json:"qty" binding:"required"`
}

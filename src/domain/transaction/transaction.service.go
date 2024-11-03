package transaction

import (
	"encoding/json"
	"errors"
	"tech-testing/src/models"
)

type TransactionService struct {
	repo *TransactionRepository
}

func NewTransactionService(repo *TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(userID string, dto *CreateTransactionDTO, wallet *models.Wallet) (string, error) {
	product, err := s.repo.FindProductByID(dto.ProductID)
	if err != nil {
		return "", errors.New("product not found")
	}

	if product.Availability < dto.Qty {
		return "", errors.New("insufficient product availability")
	}

	totalCost := product.Price * dto.Qty
	if wallet.Balance < totalCost {
		return "", errors.New("insufficient balance")
	}

	trxDetails := models.TransactionPayment{
		Qty:   dto.Qty,
		Total: totalCost,
	}

	trxDetailsJSON, err := json.Marshal(trxDetails)
	if err != nil {
		return "", errors.New("failed to marshal transaction details")
	}

	transaction := &models.Transaction{
		ProductID:       &dto.ProductID,
		WalletID:        wallet.ID,
		UserID:          userID,
		TransactionType: "PAYMENT",
		Trx:             trxDetailsJSON,
	}

	err = s.repo.ExecuteTransaction(transaction, product.ID, wallet.ID, dto.Qty, totalCost)
	if err != nil {
		return "", err
	}

	return "Transaction created successfully", nil
}

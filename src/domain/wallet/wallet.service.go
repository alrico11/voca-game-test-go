// service.go
package wallet

import (
	"encoding/json"
	"errors"
	"tech-testing/src/models"
)

type WalletService struct {
	repo *WalletRepository
}

func NewWalletService(repo *WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CheckBalance(userID string) (*models.Wallet, error) {
	return s.repo.FindByUserID(userID)
}

func (s *WalletService) Deposit(userID string, dto *WalletDepositDTO) (*models.Wallet, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	balanceAfter := wallet.Balance + dto.Balance

	trxDetails := models.TransactionDetails{
		Amount:        dto.Balance,
		BalanceBefore: wallet.Balance,
		BalanceAfter:  balanceAfter,
	}
	trxDetailsJSON, err := json.Marshal(trxDetails)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		UserID:          userID,
		WalletID:        wallet.ID,
		TransactionType: "DEPOSIT",
		Trx:             trxDetailsJSON,
	}

	// Save transaction
	if err = s.repo.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	err = s.repo.UpdateBalance(wallet, balanceAfter)
	return wallet, err
}

func (s *WalletService) Withdrawal(userID string, dto *WalletWithdrawalDTO) (*models.Wallet, error) {
	wallet, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if wallet.Balance < dto.Withdrawal {
		return nil, errors.New("insufficient balance")
	}

	balanceAfter := wallet.Balance - dto.Withdrawal

	trxDetails := models.TransactionDetails{
		Amount:        dto.Withdrawal,
		BalanceBefore: wallet.Balance,
		BalanceAfter:  balanceAfter,
	}
	trxDetailsJSON, err := json.Marshal(trxDetails)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		UserID:          userID,
		WalletID:        wallet.ID,
		TransactionType: "WITHDRAWAL",
		Trx:             trxDetailsJSON,
	}

	if err = s.repo.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	err = s.repo.UpdateBalance(wallet, balanceAfter)
	return wallet, err
}

func (s *WalletService) GetTransactionHistory(userID string, query *HistoryWalletQueryDTO) ([]models.Transaction, int64, error) {
	return s.repo.GetTransactionHistory(userID, query)
}

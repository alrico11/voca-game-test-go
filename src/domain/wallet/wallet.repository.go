// repository.go
package wallet

import (
	"tech-testing/src/models"
	"time"

	"gorm.io/gorm"
)

// WalletRepository defines the repository for wallet operations
type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) FindByUserID(userID string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) Create(wallet *models.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepository) UpdateBalance(wallet *models.Wallet, balance int) error {
	wallet.Balance = balance
	return r.db.Save(wallet).Error
}

func (r *WalletRepository) CreateTransaction(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *WalletRepository) GetTransactionHistory(userID string, query *HistoryWalletQueryDTO) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var count int64

	where := r.db.Model(&models.Transaction{}).Where("user_id = ? AND deleted_at IS NULL", userID)
	if query.Filter != "" {
		where = where.Where("transaction_type = ?", query.Filter)
	}
	if query.StartDate != "" && query.EndDate != "" {
		startDate, _ := time.Parse("2006-01-02", query.StartDate)
		endDate, _ := time.Parse("2006-01-02", query.EndDate)
		where = where.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	where.Count(&count)
	if err := where.Order(query.SortBy + " " + query.Sort).
		Limit(query.Limit).
		Offset((query.Page - 1) * query.Limit).
		Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, count, nil
}

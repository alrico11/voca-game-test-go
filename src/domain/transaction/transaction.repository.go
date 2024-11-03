package transaction

import (
	"tech-testing/src/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) FindProductByID(productID string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND deleted_at IS NULL", productID).First(&product).Error
	return &product, err
}

func (r *TransactionRepository) ExecuteTransaction(transaction *models.Transaction, productID, walletID string, qty, totalCost int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Product{}).
			Where("id = ?", productID).
			UpdateColumn("availability", gorm.Expr("availability - ?", qty)).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Wallet{}).
			Where("id = ?", walletID).
			UpdateColumn("balance", gorm.Expr("balance - ?", totalCost)).Error; err != nil {
			return err
		}

		return nil
	})
}

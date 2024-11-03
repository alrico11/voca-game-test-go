package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	PAYMENT    TransactionType = "PAYMENT"
	DEPOSIT    TransactionType = "DEPOSIT"
	WITHDRAWAL TransactionType = "WITHDRAWAL"
)

type TransactionDetails struct {
	Amount        int `json:"amount"`
	BalanceBefore int `json:"balanceBefore"`
	BalanceAfter  int `json:"balanceAfter"`
}

type TransactionPayment struct {
	Qty   int `json:"qty"`
	Total int `json:"total"`
}

type Transaction struct {
	ID              string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID          string          `gorm:"type:uuid"`
	WalletID        string          `gorm:"type:uuid"`
	ProductID       *string         `gorm:"type:uuid"`
	TransactionType TransactionType `gorm:"type:varchar(50)"`
	Trx             json.RawMessage `gorm:"type:json"`
	CreatedAt       time.Time       `gorm:"column:created_at;default:now()"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt       gorm.DeletedAt  `gorm:"column:deleted_at"`

	User    User     `gorm:"foreignKey:UserID"`
	Wallet  Wallet   `gorm:"foreignKey:WalletID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

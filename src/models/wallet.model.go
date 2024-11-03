package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID           string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Balance      int            `gorm:"default:0"`
	UserID       string         `gorm:"type:uuid;unique"`
	CreatedAt    time.Time      `gorm:"column:created_at;default:now()"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Transactions []Transaction  `gorm:"foreignKey:WalletID"`
	User         []User         `gorm:"foreignKey:WalletID"`
}

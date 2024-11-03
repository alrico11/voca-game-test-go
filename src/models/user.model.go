package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string
	Email        string `gorm:"uniqueIndex"`
	Password     string
	WalletID     *string        `gorm:"type:uuid"`
	CreatedAt    time.Time      `gorm:"column:created_at;default:now()"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Wallet       *Wallet        `gorm:"foreignKey:WalletID"`
	Transactions []Transaction  `gorm:"foreignKey:UserID"`
	UserTokens   []UserToken    `gorm:"foreignKey:UserID"`
}

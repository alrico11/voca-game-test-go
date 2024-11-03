package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string
	Description  *string
	Price        int
	Availability int
	CreatedAt    time.Time      `gorm:"column:created_at;default:now()"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
	Transactions []Transaction  `gorm:"foreignKey:ProductID"`
}

package models

import "time"

type UserToken struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    string `gorm:"type:uuid"`
	Token     string
	CreatedAt time.Time `gorm:"column:created_at;default:now()"`
	User      User      `gorm:"foreignKey:UserID"`
}

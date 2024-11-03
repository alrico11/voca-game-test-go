package auth

import (
	"tech-testing/src/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (repo *authRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

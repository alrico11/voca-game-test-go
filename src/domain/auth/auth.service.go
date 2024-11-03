package auth

import (
	"errors"
	"tech-testing/src/models"
	"tech-testing/src/utils"
)

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) Register(user *models.User) error {
	existingUser, _ := service.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return service.repo.Create(user)
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, err := service.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

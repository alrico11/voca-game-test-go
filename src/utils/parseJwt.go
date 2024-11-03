package utils

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(userID string) (string, error) {
	secret := os.Getenv("USER_JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT secret not set")
	}

	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

var (
	jwtSecret []byte
	once      sync.Once
)

func loadJWTSecret() {
	once.Do(func() {
		secret := os.Getenv("USER_JWT_SECRET")
		if secret == "" {
			panic("USER_JWT_SECRET environment variable not set")
		}
		jwtSecret = []byte(secret)
	})
}

func ParseTokenAndGetUserID(tokenStr string) (string, error) {
	loadJWTSecret()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(string)
		if !ok {
			return "", errors.New("userID not found in token")
		}
		return userID, nil
	}
	return "", errors.New("invalid token")
}

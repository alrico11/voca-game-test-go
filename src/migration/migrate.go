package migration

import (
	"log"
	"tech-testing/src/config"
	"tech-testing/src/models"
)

func RunMigrations() {
	db := config.DB

	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.UserToken{},
		&models.Wallet{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully!")
}

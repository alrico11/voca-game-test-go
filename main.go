package main

import (
	"log"
	"os"

	"tech-testing/src/config"
	"tech-testing/src/domain/auth"
	"tech-testing/src/domain/product"
	"tech-testing/src/domain/transaction"
	"tech-testing/src/domain/wallet"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	// migration.RunMigrations()

	router := gin.Default()

	auth.SetupAuthModule(config.DB, router)
	transaction.SetupTransactionModule(config.DB, router)
	product.SetupProductModule(config.DB, router)
	wallet.SetupWalletModule(config.DB, router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

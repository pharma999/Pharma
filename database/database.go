package database

import (
	"demo/config"
	"demo/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	config.LoadConfig()

	// Postgres DSN
	dsn := fmt.Sprintf(
		"host=%s user =%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	fmt.Println("✅ PostgreSQL connected!")

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}, &models.LoginPhone{}); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	DB = db
}

package database

import (
	"context"
	"demo/config"
	"demo/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MongoClient *mongo.Client
var MongoDB *mongo.Database

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
	if err := db.AutoMigrate(
		&models.User{},
		&models.LoginPhone{},
		&models.UserDetail{},
		// &models.UserReport{},
		// &models.FamilyReport{},

	); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}
	DB = db
}

func ConnectMongoDB() {
	config.LoadConfig()

	// Build URI from environment variables
	mongoHost := config.GetEnv("MONGO_HOST")
	mongoPort := config.GetEnv("MONGO_PORT")
	mongoDBName := config.GetEnv("MONGO_DBNAME")

	uri := fmt.Sprintf("mongodb://%s:%s/%s", mongoHost, mongoPort, mongoDBName)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}
	fmt.Println("✅ MongoDB connected!")

	MongoClient = client
	fmt.Println("✅ MongoDB connected! URI:", uri)
	MongoDB = client.Database("healthcare")
}

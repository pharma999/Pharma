package controller

import (
	"context"
	"demo/database"
	"demo/helper"
	"demo/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserDetail(c *gin.Context) {
	var userDetail models.UserDetail
	var loginPhone models.LoginPhone

	// Bind JSON (UUIDs are NOT expected from client)
	if err := c.ShouldBindJSON(&userDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Find loginPhone by phone number
	if err := database.DB.Where("phone_number = ?", userDetail.PhoneNumber).First(&loginPhone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number not found"})
		return
	}

	fmt.Println("Login Phone ID:", loginPhone.ID)

	// --- Generate UUIDs ---
	if userDetail.ID == uuid.Nil {
		userDetail.ID = uuid.New()
	}
	userDetail.UserId = loginPhone.ID

	// Address UUIDs
	if userDetail.Address1.UserId == uuid.Nil {
		userDetail.Address1.UserId = uuid.New()
	}
	userDetail.Address1.UserId = loginPhone.ID

	if userDetail.Address2.UserId == uuid.Nil {
		userDetail.Address2.UserId = uuid.New()
	}
	userDetail.Address2.UserId = loginPhone.ID

	// Save into DB
	if err := database.DB.Create(&userDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user detail", "details": err.Error()})
		return
	}

	helper.SucessResponse(c, gin.H{
		"message": "User detail created successfully",
		"data":    userDetail,
	})
}

func CreateUserDetailMongo(c *gin.Context) {
	var userDetail models.UserDetailRequestMongo
	var loginPhone models.LoginPhone

	// Bind JSON
	if err := c.ShouldBindJSON(&userDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Check phone number exists in SQL DB
	if err := database.DB.Where("phone_number = ?", userDetail.PhoneNumber).First(&loginPhone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number not found"})
		return
	}

	// Generate MongoDB IDs
	userDetail.ID = loginPhone.ID
	userDetail.UserId = loginPhone.ID
	userDetail.Address1.UserId = userDetail.UserId
	userDetail.Address2.UserId = userDetail.UserId

	userDetail.CreatedAt = time.Now()
	userDetail.UpdatedAt = time.Now()
	userDetail.DeletedAt = nil

	fmt.Println("Prepared UserDetail for MongoDB:", userDetail)

	// MongoDB collection
	collection := database.MongoDB.Collection("user_details")

	// Check duplicate by phone number
	filter := bson.M{"phone_number": userDetail.PhoneNumber}
	var existingUser models.UserDetailRequestMongo
	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		// User already exists
		c.JSON(http.StatusConflict, gin.H{"error": "User with this phone number already exists"})
		return // <-- IMPORTANT: stop execution
	} else if err != mongo.ErrNoDocuments {
		// Some other error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		return
	}

	// Insert new user
	_, err = collection.InsertOne(context.TODO(), userDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user detail", "details": err.Error()})
		return
	}

	helper.SucessResponse(c, gin.H{
		"message": "User detail created successfully",
		"result":  userDetail,
	})
}

// func CreateUserDetailMongo(c *gin.Context) {
// 	var userDetail models.UserDetailRequestMongo
// 	var loginPhone models.LoginPhone

// 	// Bind JSON (UUIDs are NOT expected from client)
// 	if err := c.ShouldBindJSON(&userDetail); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
// 		return
// 	}
// 	if err := database.DB.Where("phone_number = ?", userDetail.PhoneNumber).First(&loginPhone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number not found"})
// 		return
// 	}

// 	fmt.Println("Login Phone ID:", loginPhone.ID)
// 	// Map to MongoDB model

// 	// --- Generate UUIDs ---
// 	if userDetail.ID == uuid.Nil {
// 		userDetail.ID = uuid.New()
// 	}
// 	userDetail.UserId = loginPhone.ID

// 	// Address UUIDs
// 	if userDetail.Address1.UserId == uuid.Nil {
// 		userDetail.Address1.UserId = uuid.New()
// 	}
// 	userDetail.Address1.UserId = loginPhone.ID

// 	if userDetail.Address2.UserId == uuid.Nil {
// 		userDetail.Address2.UserId = uuid.New()
// 	}
// 	userDetail.Address2.UserId = loginPhone.ID

// 	// Map to MongoDB model

// 	mongoUserDetail := models.UserDetailRequestMongo{
// 		UserId:        userDetail.UserId,
// 		Name:          userDetail.Name,
// 		Email:         userDetail.Email,
// 		Gender:        userDetail.Gender,
// 		PhoneNumber:   userDetail.PhoneNumber,
// 		Image:         userDetail.Image,
// 		Address1:      userDetail.Address1,
// 		Address2:      userDetail.Address2,
// 		Status:        userDetail.Status,
// 		BlockStatus:   userDetail.BlockStatus,
// 		UserService:   userDetail.UserService,
// 		ServiceStatus: userDetail.ServiceStatus,
// 	}

// 	// Save to MongoDB
// 	collection := database.MongoDB.Collection("user_details")
// 	// check if a user with same phone number already exists
// 	filter := bson.M{"phone_number": userDetail.PhoneNumber}
// 	var existingUser models.UserDetailRequestMongo
// 	err := collection.FindOne(c, filter).Decode(&existingUser)

// 	if err == nil {
// 		// User with same phone number exists
// 		c.JSON(http.StatusConflict, gin.H{"error": "User with this phone number already exists"})
// 	} else if err != mongo.ErrNoDocuments {
// 		// Some other error occurred during the query
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
// 		return
// 	}

// 	mongoUserDetail.CreatedAt = time.Now()
// 	mongoUserDetail.UpdatedAt = time.Now()

// 	// Insert the new user detail
// 	if _, err := collection.InsertOne(c, mongoUserDetail); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user detail", "details": err.Error()})
// 		return
// 	}

// 	helper.SucessResponse(c, gin.H{
// 		"message": "User detail created successfully",
// 		"result":  mongoUserDetail,
// 	})

// }

package models

import (
	"demo/enum"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Embedded address
type Address struct {
	UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	HouseNumber string    `json:"house_number" gorm:"not null"`
	Street      string    `json:"street" gorm:"not null"`
	LandMark    string    `json:"landmark" gorm:"not null"`
	PinCode     string    `json:"pin_code" gorm:"not null"`
	Latitude    string    `json:"latitude" gorm:"not null"`
	Longitude   string    `json:"longitude" gorm:"not null"`
	IsPrimary   bool      `json:"is_primary" gorm:"default:false"`
}

// UserDetail
type UserDetail struct {
	BaseModel
	UserId      uuid.UUID   `json:"user_id" gorm:"type:uuid;uniqueIndex;not null"` // UNIQUE
	Name        string      `json:"name" gorm:"not null"`
	Email       string      `json:"email" gorm:"not null"`
	Gender      enum.Gender `json:"gender" gorm:"not null;check:gender IN ('MALE','FEMALE','OTHER')"`
	PhoneNumber string      `json:"phone_number" gorm:"not null"`
	Image       string      `json:"image_url"`
	Address1    Address     `json:"address_1" gorm:"embedded;embeddedPrefix:address1_"`
	Address2    Address     `json:"address_2" gorm:"embedded;embeddedPrefix:address2_"`
	// UserReport    []UserReport           `json:"user_report" gorm:"foreignKey:UserId;references:UserId"`
	// FamilyReport  []FamilyReport         `json:"family_report" gorm:"foreignKey:UserId;references:UserId"`
	Status        enum.UserStatus        `json:"status" gorm:"not null;default:'ACTIVE'"`
	BlockStatus   enum.BlockStatus       `json:"block_status" gorm:"not null;default:'UNBLOCKED'"`
	UserService   enum.UserServiceStatus `json:"user_service" gorm:"not null;default:'UNSUBSCRIBED'"`
	ServiceStatus enum.ServiceStatus     `json:"service_status" gorm:"not null;default:'NEW'"`
}

// BaseModelMongo for MongoDB
type BaseModelMongo struct {
	ID        uuid.UUID  `bson:"_id,omitempty" json:"id"`                          // MongoDB primary key
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`                     // Creation timestamp
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`                     // Last update timestamp
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"` // Optional soft delete
}

// AddressMongo is an embedded address struct for MongoDB
type AddressMongo struct {
	UserId      uuid.UUID `bson:"user_id" json:"user_id,omitempty"` // Reference to user
	HouseNumber string    `bson:"house_number" json:"house_number"`
	Street      string    `bson:"street" json:"street"`
	LandMark    string    `bson:"landmark" json:"landmark"`
	PinCode     string    `bson:"pin_code" json:"pin_code"`
	Latitude    string    `bson:"latitude" json:"latitude"`
	Longitude   string    `bson:"longitude" json:"longitude"`
	IsPrimary   bool      `bson:"is_primary" json:"is_primary"`
}

type UserDetailRequestMongo struct {
	BaseModelMongo
	UserId        uuid.UUID              `json:"user_id" bson:"user_id,omitempty"`
	Name          string                 `json:"name" bson:"name, omitempty"`
	Email         string                 `json:"email" bson:"email,omitempty"`
	Gender        enum.Gender            `json:"gender" bson:"gender,omitempty"`
	PhoneNumber   string                 `json:"phone_number" bson:"phone_number,omitempty"`
	Image         string                 `json:"image_url" bson:"image_url,omitempty"`
	Address1      AddressMongo           `json:"address_1" bson:"address_1,omitempty"`
	Address2      AddressMongo           `json:"address_2" bson:"address_2,omitempty"`
	Status        enum.UserStatus        `json:"status" bson:"status,omitempty"`
	BlockStatus   enum.BlockStatus       `json:"block_status" bson:"block_status,omitempty"`
	UserService   enum.UserServiceStatus `json:"user_service" bson:"user_service,omitempty"`
	ServiceStatus enum.ServiceStatus     `json:"service_status" bson:"service_status,omitempty"`
}

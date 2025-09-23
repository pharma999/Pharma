package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginPhone struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PhoneNumber string    `json:"phone_number" gorm:"not null" binding:"required"`
	IsLogin     bool      `json:"is_login" gorm:"default:false"`
}

func (loginPhone *LoginPhone) BeforeCreate(tx *gorm.DB) (err error) {
	loginPhone.ID = uuid.New()
	return nil
}

type VerifyOtp struct {
	PhoneNumber    string `json:"phone_number" gorm:"not null" binding:"required"`
	VerificationId string `json:"verification_id" gorm:"not null" binding:"required"`
	Otp            string `json:"otp" gorm:"not null" binding:"required"`
}

package controller

import (
	"demo/database"
	"demo/helper"
	"demo/middleware"
	"demo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type VerifyOTPResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
	Data         struct {
		Status         string `json:"status"`
		VerificationID string `json:"verificationId"`
	} `json:"data"`
}

func LoginUser(c *gin.Context) {
	var loginPhone models.LoginPhone
	if err := c.ShouldBindJSON(&loginPhone); err != nil {
		helper.ErrorResponce(c, "Invalid Input")
		return
	}

	// ✅ Validate phone number length (10 digits)
	if len(loginPhone.PhoneNumber) != 10 {
		helper.ErrorResponce(c, "Phone number must be 10 digits")
		return
	}

	// ✅ Check if phone number already exists
	var existingUser models.LoginPhone
	if err := database.DB.Where("phone_number = ?", loginPhone.PhoneNumber).First(&existingUser).Error; err == nil {
		// Found user → don’t insert again
		fmt.Println("User already exists, skipping DB insert")
	} else {
		// Not found → create new record
		if result := database.DB.Create(&loginPhone); result.Error != nil {
			helper.ErrorResponce(c, result.Error.Error())
			return
		}
	}

	// ✅ Send OTP
	otpResponse, err := middleware.SendOTP(loginPhone.PhoneNumber)
	if err != nil {
		helper.ErrorResponce(c, "Failed to send OTP: "+err.Error())
		return
	}

	// ✅ Check if OTP actually sent
	if otpResponse.ResponseCode == 200 &&
		otpResponse.Message == "SUCCESS" &&
		otpResponse.Data.VerificationID != "" {
		helper.SucessResponse(c, gin.H{
			"message":        "OTP sent successfully",
			"phone_number":   loginPhone.PhoneNumber,
			"verificationId": otpResponse.Data.VerificationID,
			"status":         otpResponse.ResponseCode,
			"otp_message":    otpResponse.Message,
			"timeout":        otpResponse.Data.Timeout,
			"transactionId":  otpResponse.Data.TransactionID,
		})
		return
	}

	helper.ErrorResponce(c, "Invalid phone number or failed to send OTP")
}

// Verify OTP Function
func VerifyOTP(c *gin.Context) {
	var otpVerification models.VerifyOtp
	var loginPhone models.LoginPhone
	if err := c.ShouldBindJSON(&otpVerification); err != nil {
		helper.ErrorResponce(c, "Invalid Input")
		return
	}
	resp, err := middleware.VerifyOTP(otpVerification.PhoneNumber, otpVerification.VerificationId, otpVerification.Otp)
	if err != nil {
		helper.ErrorResponce(c, "Failed to verify OTP: "+err.Error())
		return
	}

	if resp.ResponseCode == 200 &&
		resp.Message == "SUCCESS" &&
		resp.Data.VerificationStatus == "VERIFICATION_COMPLETED" {

		// do IsLogin true in db
		// if err := database.DB.Model(&loginPhone).Where("phone_number = ?", otpVerification.PhoneNumber).Update("is_login", true).Error; err != nil {
		// 	helper.ErrorResponce(c, "Failed to update login status: "+err.Error())
		// 	return
		// }
		if err := database.DB.Model(&loginPhone).
			Where("phone_number = ?", otpVerification.PhoneNumber).
			Update("is_login", true).Error; err != nil {
			helper.ErrorResponce(c, "Failed to update login status: "+err.Error())
			return
		}

		// Generate JWT Token
		token, err := helper.GenerateJWT(otpVerification.PhoneNumber)
		if err != nil {
			helper.ErrorResponce(c, "Failed to generate token: "+err.Error())
			return
		}
		helper.SucessResponse(c, gin.H{
			"message":            "OTP verified successfully",
			"phone_number":       otpVerification.PhoneNumber,
			"verificationStatus": resp.Data.VerificationStatus,
			"verificationId":     resp.Data.VerificationID,
			"token":              token,
		})
		return
	}
	helper.ErrorResponce(c, "Invalid OTP or verification failed")
}

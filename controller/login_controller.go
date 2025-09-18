package controller

import (
	"demo/database"
	"demo/helper"
	"demo/middleware"
	"demo/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func VerifyOTP(c *gin.Context) {
	fmt.Println("VerifyOTP called")
	var otpVerification models.VerifyOtp
	if err := c.ShouldBindJSON(&otpVerification); err != nil {
		helper.ErrorResponce(c, "Invalid Input")
		return
	}

	apiURL := fmt.Sprintf(
		"https://cpaas.messagecentral.com/verification/v3/validateOtp?countryCode=91&mobileNumber=%s&verificationId=%s&customerId=%s&code=%s",
		otpVerification.PhoneNumber, otpVerification.VerificationId, "C-79074004EB4443D", otpVerification.Otp,
	)
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		helper.ErrorResponce(c, "Failed to create request")
		return
	}

	// Add auth token header
	req.Header.Add("authToken", "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJDLTc5MDc0MDA0RUI0NDQzRCIsImlhdCI6MTc1ODE4MDI4OSwiZXhwIjoxOTE1ODYwMjg5fQ.UKvPC-soOGauLX7P2kIIXUCvO6UW-MZ3yzAgz6Qm5MLzqWpcujAOnjiNvb9ZHviZpYSYBN-wFWWWmx-3TcBdyQ")

	resp, err := client.Do(req)
	if err != nil {
		helper.ErrorResponce(c, "Failed to call OTP API: "+err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		helper.ErrorResponce(c, "Failed to read API response")
		return
	}

	fmt.Printf("OTP API Status: %d\n", resp.StatusCode)
	fmt.Printf("OTP API Raw Response: %s\n", string(body))

	// Check if response is valid JSON
	if !json.Valid(body) {
		helper.ErrorResponce(c, fmt.Sprintf("OTP API returned non-JSON response: %s", string(body)))
		return
	}

	// Use a struct that matches the actual API response
	var verifyResponse struct {
		ResponseCode int    `json:"responseCode"`
		Message      string `json:"message"`
		Data         struct {
			VerificationID     int     `json:"verificationId"`
			MobileNumber       string  `json:"mobileNumber"`
			VerificationStatus string  `json:"verificationStatus"`
			ResponseCode       string  `json:"responseCode"`
			ErrorMessage       *string `json:"errorMessage"`
			TransactionID      string  `json:"transactionId"`
			AuthToken          *string `json:"authToken"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &verifyResponse); err != nil {
		helper.ErrorResponce(c, fmt.Sprintf("Failed to parse API JSON: %s", string(body)))
		return
	}

	// Check all required conditions
	if verifyResponse.ResponseCode == 200 &&
		verifyResponse.Message == "SUCCESS" &&
		verifyResponse.Data.VerificationStatus == "VERIFICATION_COMPLETED" &&
		fmt.Sprintf("%d", verifyResponse.Data.VerificationID) == otpVerification.VerificationId {

		helper.SucessResponse(c, gin.H{
			"message":            "OTP verified successfully",
			"phone_number":       otpVerification.PhoneNumber,
			"verificationStatus": verifyResponse.Data.VerificationStatus,
			"verificationId":     verifyResponse.Data.VerificationID,
		})
		return
	}

	helper.ErrorResponce(c, fmt.Sprintf("Invalid OTP or verification failed: %s", verifyResponse.Message))
}

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

// func VerifyOTP(c *gin.Context) {
// 	// fmt.Println("VerifyOTP called")
// 	var otpVerification models.VerifyOtp
// 	if err := c.ShouldBindJSON(&otpVerification); err != nil {
// 		helper.ErrorResponce(c, "Invalid Input")
// 		return
// 	}

// }

// func VerifyOTP(c *gin.Context) {
// 	var otpVerification models.VerifyOtp
// 	if err := c.ShouldBindJSON(&otpVerification); err != nil {
// 		helper.ErrorResponce(c, "Invalid Input")
// 		return
// 	}

// 	// Build API URL
// 	url := fmt.Sprintf(
// 		"https://cpaas.messagecentral.com/verification/v3/validateOtp?countryCode=91&mobileNumber=%s&verificationId=%s&customerId=C-7C9DA097870D4EC&code=%s",
// 		otpVerification.PhoneNumber,
// 		otpVerification.VerificationId,
// 		otpVerification.Otp,
// 	)

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		helper.ErrorResponce(c, "Failed to create request")
// 		return
// 	}

// 	// Add auth token header
// 	req.Header.Add("authToken", "YOUR_AUTH_TOKEN_HERE")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		helper.ErrorResponce(c, "Failed to call OTP API: "+err.Error())
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		helper.ErrorResponce(c, "Failed to read API response")
// 		return
// 	}

// 	var verifyResponse VerifyOTPResponse
// 	if err := json.Unmarshal(body, &verifyResponse); err != nil {
// 		helper.ErrorResponce(c, "Failed to parse API response")
// 		return
// 	}

// 	// Check API response
// 	if verifyResponse.ResponseCode == 200 && verifyResponse.Message == "SUCCESS" {
// 		helper.SucessResponse(c, gin.H{
// 			"message":      "OTP verified successfully",
// 			"phone_number": otpVerification.PhoneNumber,
// 		})
// 		return
// 	}

//		helper.ErrorResponce(c, "Invalid OTP or verification failed")
//	}

func VerifyOTP(c *gin.Context) {
	fmt.Println("VerifyOTP called")
}

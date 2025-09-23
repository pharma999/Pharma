package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// OTPData represents the structure of the OTP response.
type OTPData struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
	Data         struct {
		VerificationID string `json:"verificationId"`
		MobileNumber   string `json:"mobileNumber"`
		ResponseCode   string `json:"responseCode"`
		Timeout        string `json:"timeout"`
		TransactionID  string `json:"transactionId"`
	} `json:"data"`
}

type VerifyOTPData struct {
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

func SendOTP(phoneNumber string) (*OTPData, error) {
	url := fmt.Sprintf(
		"https://cpaas.messagecentral.com/verification/v3/send?countryCode=91&customerId=C-79074004EB4443D&flowType=SMS&mobileNumber=%s",
		phoneNumber,
	)

	payload := strings.NewReader("")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authToken", "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJDLTc5MDc0MDA0RUI0NDQzRCIsImlhdCI6MTc1ODE4MDI4OSwiZXhwIjoxOTE1ODYwMjg5fQ.UKvPC-soOGauLX7P2kIIXUCvO6UW-MZ3yzAgz6Qm5MLzqWpcujAOnjiNvb9ZHviZpYSYBN-wFWWWmx-3TcBdyQ")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otpData OTPData
	if err := json.Unmarshal(body, &otpData); err != nil {
		return nil, err
	}

	fmt.Println("OTP Response:", otpData)
	return &otpData, nil
}

func VerifyOTP(phoneNumber, verificationId, otp string) (*VerifyOTPData, error) {
	url := fmt.Sprintf(
		"https://cpaas.messagecentral.com/verification/v3/validateOtp?countryCode=91&mobileNumber=%s&verificationId=%s&customerId=%s&code=%s",
		phoneNumber, verificationId, "C-79074004EB4443D", otp,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("authToken", "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJDLTc5MDc0MDA0RUI0NDQzRCIsImlhdCI6MTc1ODE4MDI4OSwiZXhwIjoxOTE1ODYwMjg5fQ.UKvPC-soOGauLX7P2kIIXUCvO6UW-MZ3yzAgz6Qm5MLzqWpcujAOnjiNvb9ZHviZpYSYBN-wFWWWmx-3TcBdyQ")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var verifyResp VerifyOTPData
	if err := json.Unmarshal(body, &verifyResp); err != nil {
		return nil, err
	}
	fmt.Printf("Verify OTP Response: %+v\n", verifyResp)
	return &verifyResp, nil
}

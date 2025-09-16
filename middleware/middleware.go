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

func SendOTP(phoneNumber string) (*OTPData, error) {
	url := fmt.Sprintf(
		"https://cpaas.messagecentral.com/verification/v3/send?countryCode=91&customerId=C-7C9DA097870D4EC&flowType=SMS&mobileNumber=%s",
		phoneNumber,
	)

	payload := strings.NewReader("")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authToken", "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJDLTdDOURBMDk3ODcwRDRFQyIsImlhdCI6MTc1Nzk3Mjc0NywiZXhwIjoxOTE1NjUyNzQ3fQ.AvnYUPJgFoI9ZdutvpFSDwXlugZXjgd_gvdQdyloLJXzuQrOeKuKFjgQ3N6HSS4K7gZ6CcRfv7P1LUZrG-9lXQ")

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

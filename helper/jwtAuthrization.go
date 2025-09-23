package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("my_secret_key")

// Generate JWT Token

func GenerateJWT(phone string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"phone": phone,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)

}

func ValidateJWTToken(tokenstr string) (*jwt.Token, error) {
	return jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JwtKey, nil
	})
}

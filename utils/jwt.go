package utils

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret-key")

// Generate JWT using id, role
func GenerateToken(id, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ExtractClaims extracts email, role and checks if token is still valid
func ExtractClaims(tokenString string) (email string, role string, isValid bool, err error) {
	claims := jwt.MapClaims{}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", "", false, errors.New("invalid or expired token")
	}

	// Extract values
	emailVal, ok1 := claims["email"].(string)
	roleVal, ok2 := claims["role"].(string)
	expVal, ok3 := claims["exp"].(float64) // `exp` is a float64 UNIX timestamp

	if !ok1 || !ok2 || !ok3 {
		return "", "", false, errors.New("missing or invalid claims")
	}

	// Check expiration
	if time.Now().Unix() > int64(expVal) {
		return "", "", false, errors.New("token expired")
	}

	return emailVal, roleVal, true, nil
}


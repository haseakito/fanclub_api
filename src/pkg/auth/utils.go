package auth

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims represents the JWT claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	// Set token expiry date to 3 days
	expirationTime := time.Now().Add(72 * time.Hour)

	// Set custom claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// Generate encoded token for response
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: " + err.Error())
	}

	return tokenString, nil
}

func GenerateVerificationCode() (string, error) {
	// Array to hold random bytes
	var digits [6]byte

	// Generate random numbers securely
	if _, err := rand.Read(digits[:]); err != nil {
		return "", fmt.Errorf("failed to generate random numbers securely: %w", err)
	}

	// Construct the verification code ensuring it's always six digits
	code := ""
	for _, b := range digits {
		code += fmt.Sprintf("%d", b%10) // Convert each byte to a digit in the range 0-9
	}

	return code, nil
}

func HashPassword(password string) (string, error) {
	// Hash the password for security
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: ")
	}
	return string(hashedPassword), nil
}

func VerifyPassword(password, hashedPassword string) bool {
	// Decode the hashed password and verify
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

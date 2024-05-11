package auth

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hackgame-org/fanclub_api/api/ent"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims represents the JWT claims
type Claims struct {
	User *User `json:"user"`
	jwt.StandardClaims
}

type User struct {
	UserID        string `json:"userId"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Admin         bool   `json:"admin"`
}

func GenerateToken(user *ent.User) (string, error) {
	// Set token expiry date to 3 days
	expirationTime := time.Now().Add(72 * time.Hour)

	// Set custom claims
	claims := &Claims{
		User: &User{
			UserID:        user.ID,
			Email:         user.Email,
			EmailVerified: *user.EmailVerified,
			Admin:         false,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
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

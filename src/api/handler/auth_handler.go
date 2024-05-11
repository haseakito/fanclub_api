package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/ent/verificationtoken"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/hackgame-org/fanclub_api/pkg/auth"
	"github.com/hackgame-org/fanclub_api/pkg/email"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	db *ent.Client
}

func NewAuthHandler(db *ent.Client) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

func (h AuthHandler) Signup(c echo.Context) error {
	// Bind the signup request data to SignupRequest
	var req requests.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signup request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
	}

	// Generate email verification code
	verificationCode, err := auth.GenerateVerificationCode()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate verification code"})
	}
	// Set the verification code expiry date
	verificationExpires := time.Now().Add(24 * time.Hour)

	// Start transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to start transaction"})
	}

	// Create a new user
	user, err := tx.User.Create().
		SetName(req.Name).
		SetEmail(req.Email).
		SetPassword(hashedPassword).
		SetProfileImageURL("https://dbdehqz6rw0l.cloudfront.net/profile/default-profile.jpeg").
		Save(c.Request().Context())
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}

	// Create an email verification code
	_, err = tx.VerificationToken.Create().
		SetUser(user).
		SetEmail(user.Email).
		SetVerificationCode(verificationCode).
		SetExpiresAt(verificationExpires).
		Save(c.Request().Context())
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create email verification code"})
	}

	// Commit the transaction
	tx.Commit()

	// Send email containing the verification code
	email.SendMail(
		c.Request().Context(),
		req.Email,
		"Verify your email",
		"Enter the OTP "+verificationCode,
	)

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h AuthHandler) Signin(c echo.Context) error {
	// Bind the signup request data to SigninRequest
	var req requests.SigninRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signin request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Retrieve the user with email
	user, err := h.db.User.Query().
		Where(user.Email(req.Email)).
		First(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid email or password"})
	}

	// Verify the password
	if !auth.VerifyPassword(req.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":      "Successfully signed in.",
		"access_token": token,
	})
}

func (h AuthHandler) VerifyEmail(c echo.Context) error {
	// Bind the signup request data to VerifyEmailRequest
	var req requests.VerifyEmailRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signup request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Query the token
	token, err := h.db.VerificationToken.
		Query().
		Where(verificationtoken.Email(req.Email)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Check if the verification token has expired.
	if token.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Verification code has expired"})
	}

	// Compare the provided code with the stored verification code.
	if token.VerificationCode != req.Code {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Incorrect verification code"})
	}

	// Start transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to start transaction"})
	}

	// Update user fields
	_, err = tx.User.
		Update().
		Where(user.Email(req.Email)).
		SetEmailVerified(true).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Delete the token
	err = tx.VerificationToken.
		DeleteOne(token).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Commit the transaction
	tx.Commit()

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully verified this email"})
}

func (h AuthHandler) ResendVerifyEmail(c echo.Context) error {
	// Bind the signup request data to ResendVerifyEmailRequest
	var req requests.ResendVerifyEmailRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signup request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Query an user with the user id
	user, err := h.db.User.
		Query().
		Where(user.Email(req.Email)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Query the token
	_, err = h.db.VerificationToken.
		Delete().
		Where(verificationtoken.Email(req.Email)).
		Exec(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Generate email verification code
	verificationCode, err := auth.GenerateVerificationCode()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate verification code"})
	}
	// Set the verification code expiry date
	verificationExpires := time.Now().Add(24 * time.Hour)

	// Create an email verification code
	_, err = h.db.VerificationToken.Create().
		SetUser(user).
		SetEmail(user.Email).
		SetVerificationCode(verificationCode).
		SetExpiresAt(verificationExpires).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create password reset code"})
	}

	// Send email containing the verification code
	email.SendMail(
		c.Request().Context(),
		req.Email,
		"Reset password",
		"Enter the OTP "+verificationCode,
	)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully resend verification token"})
}

func (h AuthHandler) ResetPassword(c echo.Context) error {
	// Bind the signup request data to ResetPasswordRequest
	var req requests.ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signup request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Query an user with the user id
	user, err := h.db.User.
		Query().
		Where(user.Email(req.Email)).
		WithVerificationToken().
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// If there's already token issued, first delete it
	if user.Edges.VerificationToken != nil {
		// Delete the token
		err = h.db.VerificationToken.
			DeleteOne(user.Edges.VerificationToken).
			Exec(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	}

	// Generate email verification code
	verificationCode, err := auth.GenerateVerificationCode()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate verification code"})
	}
	// Set the verification code expiry date
	verificationExpires := time.Now().Add(24 * time.Hour)

	// Create an email verification code
	_, err = h.db.VerificationToken.Create().
		SetUser(user).
		SetEmail(user.Email).
		SetVerificationCode(verificationCode).
		SetExpiresAt(verificationExpires).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create password reset code"})
	}

	// Construct the reset password URL with properly encoded parameters
	resetURL := fmt.Sprintf("http://localhost:3000/auth/reset-password/confirm?email=%s&code=%s",
		url.QueryEscape(user.Email), url.QueryEscape(verificationCode))

	// Send email containing the verification code
	email.SendMail(
		c.Request().Context(),
		req.Email,
		"Reset password",
		"Click the link "+resetURL,
	)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully sent password reset email"})
}

func (h AuthHandler) VerifyPasswordReset(c echo.Context) error {
	// Bind the signup request data to VerifyPasswordResetRequest
	var req requests.VerifyPasswordResetRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Validate the signup request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
	}

	// Query the token
	token, err := h.db.VerificationToken.
		Query().
		Where(verificationtoken.Email(req.Email)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Check if the verification token has expired.
	if token.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Verification code has expired"})
	}

	// Compare the provided code with the stored verification code.
	if token.VerificationCode != req.Code {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Incorrect verification code"})
	}

	// Start transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to start transaction"})
	}

	// Update user fields
	_, err = tx.User.
		Update().
		Where(user.Email(req.Email)).
		SetPassword(hashedPassword).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Delete the token
	err = tx.VerificationToken.
		DeleteOne(token).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Commit the transaction
	tx.Commit()

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully reset password"})
}

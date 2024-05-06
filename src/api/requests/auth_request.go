package requests

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req SignupRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Name,
			validation.Required,
			validation.Length(1, 30),
		),
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email,
		),
		validation.Field(
			&req.Password,
			validation.Required,
			validation.Length(8, 1000),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`\d`)).Error("must contain at least one digit"),
		),
	)
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req SigninRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email,
		),
		validation.Field(
			&req.Password,
			validation.Required,
			validation.Length(8, 1000),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`\d`)).Error("must contain at least one digit"),
		),
	)
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (req VerifyEmailRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email.Error("Invalid email format"),
		),
		validation.Field(
			&req.Code,
			validation.Required.Error("The verification code is required"),
			validation.Length(6, 6).Error("The verification code must be exactly 6 characters long"),
			is.Digit.Error("The verification code must consist only of digits"),
		),
	)
}

type ResendVerifyEmailRequest struct {
	Email string `json:"email"`
}

func (req ResendVerifyEmailRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email.Error("Invalid email format"),
		),
	)
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func (req ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email.Error("Invalid email format"),
		),
	)
}

type VerifyPasswordResetRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (req VerifyPasswordResetRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email.Error("Invalid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required,
			validation.Length(8, 1000),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`\d`)).Error("must contain at least one digit"),
		),
		validation.Field(
			&req.Code,
			validation.Required.Error("The verification code is required"),
			validation.Length(6, 6).Error("The verification code must be exactly 6 characters long"),
			is.Digit.Error("The verification code must consist only of digits"),
		),
	)
}

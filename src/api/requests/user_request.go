package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ProfileUpdateRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
	Url  string `json:"url"`
}

func (req ProfileUpdateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Name,
			validation.Length(1, 100),
		),
		validation.Field(
			&req.Bio,
			validation.Length(0, 1000),
		),
		validation.Field(
			&req.Url,
			is.URL,
		),
	)
}

type AcountUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	DOB      string `json:"dob"`
}

func (req AcountUpdateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Username,
			validation.Length(1, 100),
		),
		validation.Field(
			&req.Email,
			validation.Required,
			validation.Length(1, 100),
			is.Email.Error("Invalid email format"),
		),
		validation.Field(
			&req.DOB,
			validation.Required,
		),
	)
}

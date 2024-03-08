package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserRequest struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Url      string `json:"url"`
}

func (req UserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Username,
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

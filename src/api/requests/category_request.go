package requests

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type CategoryRequest struct {
	Name string `json:"name"`
}

func (req CategoryRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Name,
			validation.Required,
			validation.Length(1, 100),
		),
	)
}

package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BillboardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (req BillboardRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Title,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(
			&req.Description,
			validation.Length(0, 1000),
		),
	)
}

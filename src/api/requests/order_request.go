package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type OrderRequest struct {
	PostID string `json:"postId"`
}

func (req OrderRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.PostID,
			validation.Required,
		),
	)
}

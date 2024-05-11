package requests

import "github.com/go-ozzo/ozzo-validation/v4"

type LikeRequest struct {
	PostID string `json:"postId"`
}

func (req LikeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.PostID,
			validation.Required,
		),
	)
}

package requests

import "github.com/go-ozzo/ozzo-validation/v4"

type LikeRequest struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (req LikeRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.UserID,
			validation.Required,
		),
		validation.Field(
			&req.PostID,
			validation.Required,
		),
	)
}

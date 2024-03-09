package requests

import "github.com/go-ozzo/ozzo-validation/v4"

type FollowRequest struct {
	FollowerID  string `json:"follower_id"`
	FollowingID string `json:"following_id"`
}

func (req FollowRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.FollowerID,
			validation.Required,
		),
		validation.Field(
			&req.FollowingID,
			validation.Required,
		),
	)
}

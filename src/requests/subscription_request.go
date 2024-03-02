package requests

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type SubscriptionRequest struct {
	Name        string  `json:"name"`
	UserID      string  `json:"user_id"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	TrialPeriod *int    `json:"trial_period_days"`
	IsArchived  bool    `json:"is_archived"`
}

func (req SubscriptionRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Name,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(
			&req.UserID,
			validation.Required,
		),
		validation.Field(
			&req.Description,
			validation.Length(0, 1000),
		),
		validation.Field(
			&req.Price,
			validation.Min(0),
		),
		validation.Field(
			&req.TrialPeriod,
			validation.Min(0),
		),
	)
}

package requests

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type PostCreateRequest struct {
	Title string `json:"title"`
}

func (req PostCreateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Title,
			validation.Required,
			validation.Length(1, 100),
		),
	)
}

type PostRequest struct {
	Title           string   `json:"title"`
	Description     *string  `json:"description"`
	Price           *int     `json:"price"`
	CategoryIDs     []string `json:"categories"`
	SubscriptionIDs []string `json:"subscriptions"`
	IsFeatured      bool     `json:"is_featured"`
	Status          bool     `json:"status"`
}

func (req PostRequest) Validate() error {
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
		validation.Field(
			&req.Price,
			validation.Min(0),
		),
		validation.Field(
			&req.CategoryIDs,
			validation.Each(validation.Length(1, 100)),
		),
		validation.Field(
			&req.SubscriptionIDs,
			validation.Each(validation.Length(1, 100)),
		),
	)
}

package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/order"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v78"
)

type WebhookHandler struct {
	db *ent.Client
}

func NewWebhookHandler(db *ent.Client) *WebhookHandler {
	return &WebhookHandler{
		db: db,
	}
}

func (h WebhookHandler) StripeWebhook(c echo.Context) error {
	// Read the request body
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Bind the event data to Stripe event
	var event stripe.Event
	if err := json.Unmarshal(body, &event); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse Stripe webhook payload: " + err.Error()})
	}

	// Process the event
	switch event.Type {
	case stripe.EventTypeAccountUpdated:
		// Bind the event data to Account
		var accountUpdate stripe.Account
		err := json.Unmarshal(event.Data.Raw, &accountUpdate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse Stripe account: " + err.Error()})
		}

		// Update role field to creator
		_, err = h.db.User.
			Update().
			Where(user.Email(accountUpdate.Email)).
			SetRole(user.RoleCreator).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update user role: " + err.Error()})
		}
	case stripe.EventTypePaymentIntentSucceeded:
		// Bind the event data to PaymentIntent
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse Stripe payment intent: " + err.Error()})
		}

		// Fetch order id from metadata
		orderID := paymentIntent.Metadata["order_id"]

		// Update status field to completed
		_, err = h.db.Order.
			UpdateOneID(orderID).
			SetStatus(order.StatusCompleted).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update order status: " + err.Error()})
		}
	case stripe.EventTypePaymentIntentProcessing:
		// Bind the event data to PaymentIntent
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse Stripe payment intent: " + err.Error()})
		}

		// Fetch order id from metadata
		orderID := paymentIntent.Metadata["order_id"]

		// Update status field to processing
		_, err = h.db.Order.
			UpdateOneID(orderID).
			SetStatus(order.StatusProcessing).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update order status: " + err.Error()})
		}
	case stripe.EventTypePaymentIntentPaymentFailed:
		// Bind the event data to PaymentIntent
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse Stripe payment intent: " + err.Error()})
		}

		// Fetch order id from metadata
		orderID := paymentIntent.Metadata["order_id"]

		// Update status field to canceled
		_, err = h.db.Order.
			UpdateOneID(orderID).
			SetStatus(order.StatusCanceled).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update order status: " + err.Error()})
		}
	}

	return c.NoContent(http.StatusOK)
}

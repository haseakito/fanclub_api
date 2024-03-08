package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/labstack/echo/v4"
	svix "github.com/svix/svix-webhooks/go"
)

type WebhookHandler struct {
	db *ent.Client
}

func NewWebhookHandler(db *ent.Client) *WebhookHandler {
	return &WebhookHandler{
		db: db,
	}
}

// ClerkEvent represents events received from Clerk with Svix.
type ClerkEvent struct {
	Payload   json.RawMessage `json:"data"`
	Object    string          `json:"object"`
	EventType string          `json:"type"`
}

// UserPayload represents the payload structure for user-related events.
type UserPayload struct {
	Birthday              string            `json:"birthday"`
	CreatedAt             int64             `json:"created_at"`
	EmailAddresses        []EmailAddress    `json:"email_addresses"`
	ExternalAccounts      []interface{}     `json:"external_accounts"` // Placeholder for external accounts
	ExternalID            string            `json:"external_id"`
	FirstName             string            `json:"first_name"`
	Gender                string            `json:"gender"`
	ID                    string            `json:"id"`
	ImageURL              string            `json:"image_url"`
	LastName              string            `json:"last_name"`
	LastSignInAt          int64             `json:"last_sign_in_at"`
	Object                string            `json:"object"`
	PasswordEnabled       bool              `json:"password_enabled"`
	PhoneNumbers          []interface{}     `json:"phone_numbers"` // Placeholder for phone numbers
	PrimaryEmailAddressID string            `json:"primary_email_address_id"`
	PrimaryPhoneNumberID  interface{}       `json:"primary_phone_number_id"`
	PrimaryWeb3WalletID   interface{}       `json:"primary_web3_wallet_id"`
	PrivateMetadata       map[string]string `json:"private_metadata"`
	ProfileImageURL       string            `json:"profile_image_url"`
	PublicMetadata        map[string]string `json:"public_metadata"`
	TwoFactorEnabled      bool              `json:"two_factor_enabled"`
	UnsafeMetadata        map[string]string `json:"unsafe_metadata"`
	UpdatedAt             int64             `json:"updated_at"`
	Username              interface{}       `json:"username"`
	Web3Wallets           []interface{}     `json:"web3_wallets"` // Placeholder for web3 wallets
}

// EmailAddress represents the structure of an email address in the payload.
type EmailAddress struct {
	EmailAddress string `json:"email_address"`
	ID           string `json:"id"`
	LinkedTo     []struct {
	} `json:"linked_to"`
	Object       string `json:"object"`
	Verification struct {
		Status   string `json:"status"`
		Strategy string `json:"strategy"`
	} `json:"verification"`
}

// UserDeletedPayload represents the payload structure for the "user.deleted" event.
type UserDeletedPayload struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}

func (h WebhookHandler) ClerkWebhook(c echo.Context) error {
	// Initialize the svix webhook listener
	wh, err := svix.NewWebhook(os.Getenv("CLERK_WEBHOOK_SECRET"))
	if err != nil {
		log.Printf("Failed to initialize svix webhhok: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Get request body
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// Verify the payload with request headers
	err = wh.Verify(body, c.Request().Header)
	if err != nil {
		log.Printf("Failed to verify the request signature: %v", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Parse the JSON data into ClerkEvent
	var evt ClerkEvent
	err = json.Unmarshal(body, &evt)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Switch statement to handle different event types
	switch evt.EventType {
	// Create event
	case "user.created":
		{
			// Bind the payload data to UserPayload
			var payload UserPayload
			if err := json.Unmarshal(evt.Payload, &payload); err != nil {
				log.Fatal("Error parsing payload:", err)
			}

			// Insert a new user
			err := h.db.User.
				Create().
				SetID(payload.ID).
				SetName(payload.FirstName + " " + payload.LastName).
				SetEmail(payload.EmailAddresses[0].EmailAddress).
				SetProfileImageURL(payload.ProfileImageURL).
				SaveX(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
	// Update event
	case "user.updated":
		{
			// Bind the payload data to UserPayload
			var payload UserPayload
			if err := json.Unmarshal(evt.Payload, &payload); err != nil {
				log.Fatal("Error parsing payload:", err)
			}

			// Update fields
			err := h.db.User.
				UpdateOneID(payload.ID).
				SetName(payload.FirstName + " " + payload.LastName).
				SetEmail(payload.EmailAddresses[0].EmailAddress).
				SetProfileImageURL(payload.ProfileImageURL).
				SaveX(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
	// Delete event
	case "user.deleted":
		{
			// Bind the payload data to UserDeletedPayload
			var payload UserDeletedPayload
			if err := json.Unmarshal(evt.Payload, &payload); err != nil {
				log.Fatal("Error parsing payload:", err)
			}

			err := h.db.User.
				DeleteOneID(payload.ID).
				Exec(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
	}

	return c.JSON(http.StatusOK, "Successfully received the webhook events fron Clerk")
}
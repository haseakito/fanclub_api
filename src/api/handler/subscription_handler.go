package handlers

import (
	"net/http"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	db *ent.Client
}

func NewSubscriptionHandler(db *ent.Client) *SubscriptionHandler {
	return &SubscriptionHandler{
		db: db,
	}
}

func (h SubscriptionHandler) CreateSubscription(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := h.db.User.Query().Where(user.ID(req.UserID)).Only(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Insert a new subscription data
	sub, err := h.db.Subscription.
		Create().
		SetUser(user).
		SetName(req.Name).
		SetDescription(*req.Description).
		SetPrice(*req.Price).
		SetTrialPeriodDays(*req.TrialPeriod).
		SetIsArchived(req.IsArchived).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, sub)
}

func (h SubscriptionHandler) GetSubscriptions(c echo.Context) error {
	// Get user id from query parameter
	userID := c.QueryParam("user_id")

	// Get subscriptions for a user with user id
	res, err := h.db.User.
		QuerySubscriptions(h.db.User.GetX(c.Request().Context(), userID)).
		All(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h SubscriptionHandler) GetSubscription(c echo.Context) error {
	// Get subscription id from request parameter
	subscriptionID := c.Param("id")

	// Query the subscription with subscription id
	res, err := h.db.Subscription.Get(c.Request().Context(), subscriptionID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h SubscriptionHandler) UpdateSubscription(c echo.Context) error {
	// Get subscription id from request parameter
	subscriptionID := c.Param("subscriptionId")

	// Bind the request data to PostRequet
	var req requests.SubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Update fields with data provided
	res, err := h.db.Subscription.
		UpdateOneID(subscriptionID).
		SetName(req.Name).
		SetDescription(*req.Description).
		SetPrice(*req.Price).
		SetTrialPeriodDays(*req.TrialPeriod).
		SetIsArchived(req.IsArchived).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (h SubscriptionHandler) DeleteSubscription(c echo.Context) error {
	// Get subscription id from request parameter
	subscriptionID := c.Param("id")

	// Delete the subscription with subscription id
	if err := h.db.Subscription.
		DeleteOneID(subscriptionID).
		Exec(c.Request().Context()); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete subscription")
	}

	return c.JSON(http.StatusOK, "Successfully delete the subscription")
}

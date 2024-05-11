package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/order"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/hackgame-org/fanclub_api/internal/redis"
	"github.com/hackgame-org/fanclub_api/pkg/cache"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

type OrderHandler struct {
	db    *ent.Client
	cache *cache.Cache[[]*ent.Order]
}

func NewOrderHandler(db *ent.Client, rdb redis.Client) *OrderHandler {
	return &OrderHandler{
		db: db,
		cache: cache.NewCache[[]*ent.Order](
			rdb,
			time.Minute,
		),
	}
}

func (h OrderHandler) CreateOrder(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Bind the request data to OrderRequest
	var req requests.OrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Fetch post with post id
	post, err := h.db.Post.
		Query().
		Where(post.ID(req.PostID)).
		WithUser().
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Calculate the fee for our platform
	fee := int64(0.15 * float64(post.Price))
	netAmount := int64(post.Price - fee)

	// Start a new transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Create a new order
	order, err := tx.Order.
		Create().
		SetPostID(req.PostID).
		SetUserID(userID).
		SetAmount(netAmount).
		Save(c.Request().Context())
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"Failed to create an order": err.Error()})
	}

	// Instantiate a new Stripe session parameter
	param := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Quantity: stripe.Int64(1),
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name:        stripe.String(post.Title),
						Description: stripe.String(post.Description),
					},
				},
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			ApplicationFeeAmount: stripe.Int64(fee),
			TransferData: &stripe.CheckoutSessionPaymentIntentDataTransferDataParams{
				Destination: stripe.String(post.Edges.User.StripeAccountID),
			},
		},
		Metadata: map[string]string{
			"order_id": order.ID,
		},
		SuccessURL: stripe.String("http://localhost:3000/posts/" + req.PostID + "?checkout=success"),
		CancelURL:  stripe.String("http://localhost:3000/posts/" + req.PostID + "?checkout=cancel"),
	}

	// Create a new Stripe checkout session
	res, err := session.New(param)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"Failed to create a Stripe checkout session": err.Error()})
	}

	// Commit the transaction
	tx.Commit()

	return c.JSON(http.StatusOK, res)
}

func (h OrderHandler) GetOrders(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Get limit from query parameter
	limitStr := c.QueryParam("limit")
	// Convert the limit from string to int
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to parse pagination limit"})
	}

	// Get offset from query parameter
	offsetStr := c.QueryParam("offset")
	// Convert the offset from string to int
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to parse pagination offset"})
	}

	// Define a key for caching that includes the pagination parameters
	cacheKey := fmt.Sprintf("orders-limit-%d-offset-%d", limit, offset)

	// Query orders with limit and offset
	orders, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.Order, error) {
		return h.db.Order.Query().
			Where(order.HasUserWith(user.ID(userID))).
			Offset(offset).
			Limit(limit).
			All(c.Request().Context())
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch orders: " + err.Error()})
	}

	return c.JSON(http.StatusOK, orders)
}

func (h OrderHandler) GetOrder(c echo.Context) error {
	// Get order id from request parameter
	orderID := c.Param("id")

	// Fetch order with order id
	order, err := h.db.Order.Query().
		Where(order.ID(orderID)).
		WithPost().
		WithUser().
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, order)
}

package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db *ent.Client
}

func NewUserHandler(db *ent.Client) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (h UserHandler) GetUsers(c echo.Context) error {
	return nil
}

// TODO: return user info with subscriptions
func (h UserHandler) GetUser(c echo.Context) error {
	return nil
}

func (h UserHandler) GetPostsByUserID(c echo.Context) error {
	// Get user id from request parameter
	userID := c.Param("id")

	// Get limit from query parameter
	limitStr := c.QueryParam("limit")
	// Convert the limit from string to int
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse pagination limit")
	}

	// Get offset from query parameter
	offsetStr := c.QueryParam("offset")
	// Convert the offset from string to int
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse pagination offset")
	}

	// Query posts by user id with limit and offset
	res, err := h.db.Post.
		Query().
		Where(post.UserID(userID)).
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, res)
}
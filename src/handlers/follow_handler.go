package handlers

import (
	"net/http"

	"github.com/hackgame-org/fanclub_api/ent"	
	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/labstack/echo/v4"
)

type FollowHandler struct {
	db *ent.Client
}

func NewFollowHandler(db *ent.Client) *FollowHandler {
	return &FollowHandler{
		db: db,
	}
}

func (h FollowHandler) CreateFollow(c echo.Context) error {
	// Bind the request data to FollowRequet
	var req requests.FollowRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Insert a new follow with follower id and following user id
	_, err := h.db.User.
		UpdateOneID(req.FollowerID).
		AddFollowers(h.db.User.GetX(c.Request().Context(), req.FollowingID)).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "Successfully followed the user")
}

func (h FollowHandler) DeleteFollow(c echo.Context) error {
	// Bind the request data to FollowRequet
	var req requests.FollowRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Delete the follow with follower id and following user id
	_, err := h.db.User.
		UpdateOneID(req.FollowerID).
		RemoveFollowers(h.db.User.GetX(c.Request().Context(), req.FollowingID)).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "Successfully unfollowed the user")
}

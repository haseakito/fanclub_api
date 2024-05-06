package handlers

import (
	"net/http"
	
	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/like"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	db *ent.Client
}

func NewLikeHandler(db *ent.Client) *LikeHandler {
	return &LikeHandler{
		db: db,
	}
}

// TODO: Add middleware and check if the user id matches
func (h LikeHandler) CreateLike(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Query the like with user id and post id
	exists, err := h.db.Like.
		Query().
		Where(
			like.HasUserWith(user.ID(req.UserID)),
			like.HasPostWith(post.ID(req.PostID)),
		).
		Exist(c.Request().Context())
	// If the like failed, then throw an error
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query like"+err.Error())
	}

	// If the post is already liked, then throw an error
	if exists {
		return c.String(http.StatusConflict, "Already liked the post")
	}

	// Create a like to the post
	_, err = h.db.Like.Create().
		SetUserID(req.UserID).
		SetPostID(req.PostID).
		Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to like the post"+err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

// TODO: Add middleware and check if the user id matches
func (h LikeHandler) DeleteLike(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Delete the like with post id and user id
	_, err := h.db.Like.
		Delete().
		Where(
			like.HasUserWith(user.ID(req.UserID)),
			like.HasPostWith(post.ID(req.PostID)),
		).
		Exec(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete like"+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/like"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/hackgame-org/fanclub_api/internal/redis"
	"github.com/hackgame-org/fanclub_api/pkg/cache"
	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	db    *ent.Client
	cache *cache.Cache[[]*ent.Post]
}

func NewLikeHandler(db *ent.Client, rdb redis.Client) *LikeHandler {
	return &LikeHandler{
		db: db,
		cache: cache.NewCache[[]*ent.Post](
			rdb,
			time.Minute,
		),
	}
}

func (h LikeHandler) CreateLike(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Bind the request data to LikeRequet
	var req requests.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Query the like with user id and post id
	exists, err := h.db.Like.
		Query().
		Where(
			like.HasUserWith(user.ID(userID)),
			like.HasPostWith(post.ID(req.PostID)),
		).
		Exist(c.Request().Context())
	// If the like failed, then throw an error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to query like" + err.Error()})
	}

	// If the post is already liked, then throw an error
	if exists {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Already liked the post"})
	}

	// Create a like to the post
	_, err = h.db.Like.Create().
		SetUserID(userID).
		SetPostID(req.PostID).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to like the post: " + err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (h LikeHandler) DeleteLike(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Bind the request data to LikeRequet
	var req requests.LikeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Delete the like with post id and user id
	_, err := h.db.Like.
		Delete().
		Where(
			like.HasUserWith(user.ID(userID)),
			like.HasPostWith(post.ID(req.PostID)),
		).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete like: " + err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (h LikeHandler) GetPostsByLike(c echo.Context) error {
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

	// Define a key for caching that includes the user ID, limit, and offset
	cacheKey := fmt.Sprintf("user-%s-posts-liked", userID)

	res, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.Post, error) {
		// Fetch posts the user has liked
		return h.db.Like.
			Query().
			Where(like.HasUserWith(user.ID(userID))).
			QueryPost().
			Limit(limit).
			Offset(offset).
			All(ctx)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch posts: " + err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

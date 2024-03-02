package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	db *ent.Client
}

func NewPostHandler(db *ent.Client) *PostHandler {
	return &PostHandler{
		db: db,
	}
}

func (h PostHandler) CreatePost(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.PostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Insert a new post data
	res, err := h.db.Post.
		Create().
		SetTitle(req.Title).
		SetUserID(req.UserID).
		Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h PostHandler) GetPosts(c echo.Context) error {
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

	// Query posts with limit and offset
	res, err := h.db.Post.
		Query().
		WithCategories().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h PostHandler) GetPostByID(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Query the post with post id
	res, err := h.db.Post.
		Query().
		Where(post.ID(postUUID)).
		WithCategories().
		Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h PostHandler) UpdatePost(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Bind the request data to PostRequet
	var req requests.PostRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	// Start a new transaction
	tx, err := h.db.Tx(context.Background())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Update fields with data provided
	res, err := tx.Post.
		UpdateOneID(postUUID).
		SetTitle(req.Title).
		SetUserID(req.UserID).
		SetDescription(*req.Description).
		SetPrice(*req.Price).
		SetIsFeatured(req.IsFeatured).
		SetStatus(req.Status).
		Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Iterate through categories to associate the post with categories
	for _, categoryID := range req.CategoryIDs {
		// Parse category ID string into UUID
		categoryUUID, err := uuid.Parse(categoryID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Query the category with category id
		category, err := tx.Category.Get(context.Background(), categoryUUID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// Update the category field for the post
		if _, err = res.Update().
			AddCategories(category).
			Save(context.Background()); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}



func (h PostHandler) DeletePost(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Delete the post with post id
	if err := h.db.Post.
		DeleteOneID(postUUID).
		Exec(context.Background()); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete post")
	}

	return c.JSON(http.StatusOK, "Successfully delete the post")
}

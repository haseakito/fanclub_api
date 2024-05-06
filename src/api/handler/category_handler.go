package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	db *ent.Client
}

func NewCategoryHandler(db *ent.Client) *CategoryHandler {
	return &CategoryHandler{
		db: db,
	}
}

func (h CategoryHandler) CreateCategories(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.CategoryRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return err
	}

	// Insert a new category
	res, err := h.db.Category.
		Create().
		SetName(req.Name).
		Save(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create a new category: "+err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h CategoryHandler) GetCategories(c echo.Context) error {
	// Query all categories
	res, err := h.db.Category.
		Query().
		All(context.Background())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h CategoryHandler) GetPostsByCategoryID(c echo.Context) error {
	// Get category id from request parameter
	categoryID := c.Param("id")

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

	// Query posts by category id with limit and offset
	res, err := h.db.Category.
		QueryPosts(h.db.Category.GetX(context.Background(), categoryID)).
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h CategoryHandler) DeleteCategory(c echo.Context) error {
	// Get category id from request
	categoryID := c.Param("id")

	// Delete the category with category id
	if err := h.db.Category.
		DeleteOneID(categoryID).
		Exec(context.Background()); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete category")
	}

	return c.JSON(http.StatusOK, "Successfully delete the category")
}

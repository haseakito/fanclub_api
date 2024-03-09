package handlers

import (
	"net/http"
	"strconv"

	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/user"
	"github.com/hackgame-org/fanclub_api/requests"
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

func (h UserHandler) FollowUser(c echo.Context) error {
	// Get user id from request parameter
	userID := c.Param("id")


	h.db.User.
		UpdateOneID(userID).
		Save(c.Request().Context())

	return nil
}

func (h UserHandler) GetUsers(c echo.Context) error {
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
	users, err := h.db.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users)
	}

	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) GetUser(c echo.Context) error {
	// Get user id from request parameter
	userID := c.Param("id")

	// Query user with user id
	user, err := h.db.User.
		Query().
		Where(user.ID(userID)).
		WithPosts(func(pq *ent.PostQuery) { pq.Limit(6) }).
		WithSubscriptions().
		Only(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user)
	}

	// Query followers count
	followersCount, err := h.db.User.
		QueryFollowers(user).
		Count(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user)
	}

	// Query followers count
	followingCount, err := h.db.User.
		QueryFollowing(user).
		Count(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user)
	}

	// Create a response structure that includes the user and the following and followers count
	type UserResponse struct {
		User           *ent.User `json:"user"`
		FollowingCount int       `json:"following"`
		FollowersCount int       `json:"followers"`
	}

	// Create the response object
	res := UserResponse{
		User:           user,
		FollowingCount: followingCount,
		FollowersCount: followersCount,
	}

	return c.JSON(http.StatusOK, res)
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
	posts, err := h.db.User.
		QueryPosts(h.db.User.GetX(c.Request().Context(), userID)).
		Limit(limit).
		Offset(offset).
		All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, posts)
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	// Get user id from request parameter
	userID := c.Param("id")

	// Bind the request data to UserRequest
	var req requests.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Update fields
	user, err := h.db.User.
		UpdateOneID(userID).
		SetUsername(req.Username).
		SetBio(req.Bio).
		SetURL(req.Url).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

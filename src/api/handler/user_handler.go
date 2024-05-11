package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hackgame-org/fanclub_api/api/ent"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/hackgame-org/fanclub_api/internal/redis"

	"github.com/hackgame-org/fanclub_api/pkg/cache"
	"github.com/hackgame-org/fanclub_api/pkg/storage"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db    *ent.Client
	cache *cache.Cache[[]*ent.User]
}

func NewUserHandler(db *ent.Client, rdb redis.Client) *UserHandler {
	return &UserHandler{
		db: db,
		cache: cache.NewCache[[]*ent.User](
			rdb,
			time.Minute,
		),
	}
}

func (h UserHandler) UploadProfilePicture(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Query an user with the user id
	user, err := h.db.User.Get(c.Request().Context(), userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// If there's already an profile picture, delete it first before uploading a new one
	if user.ProfileImageURL != "" {
		storage.DeleteFile(c.Request().Context(), user.ProfileImageURL)
	}

	// Read the form file
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to read file: " + err.Error()})
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to open file: " + err.Error()})
	}

	// Construct a unique key for the S3 object
	key := "profiles/" + userID + "-" + file.Filename

	// Upload file to storage
	loc, err := storage.UploadFile(c.Request().Context(), src, file.Header.Get("Content-Type"), key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload file: " + err.Error()})
	}

	// Update profile image url field
	_, err = h.db.User.UpdateOneID(userID).
		SetProfileImageURL(loc).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update the user profile: " + err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (h UserHandler) FollowUser(c echo.Context) error {
	// Get the id of a user to be followed from request parameter
	followeeID := c.Param("id")

	// Get the user id from context
	userID := c.Get("userID").(string)

	// Create the following relationship
	err := h.db.User.UpdateOneID(userID).
		AddFollowerIDs(followeeID).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to follow user: " + err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (h UserHandler) UnfollowUser(c echo.Context) error {
	// Get the id of a user to be unfollowed from request parameter
	followeeID := c.Param("id")

	// Get the user id from context
	userID := c.Get("userID").(string)

	// Destory the following relationship
	err := h.db.User.UpdateOneID(userID).
		RemoveFollowingIDs(followeeID).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to unfollow user: " + err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func (h UserHandler) GetUsers(c echo.Context) error {
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
	cacheKey := fmt.Sprintf("users-limit-%d-offset-%d", limit, offset)

	// Query posts with limit and offset
	users, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.User, error) {
		return h.db.User.
			Query().
			Limit(limit).
			Offset(offset).
			All(ctx)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users)
	}

	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) GetUser(c echo.Context) error {
	// Get user id from request parameter
	userID := c.Param("id")

	// Fetch user from cache first, then from database
	user, err := h.db.User.Get(c.Request().Context(), userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
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

	// Create the response object
	res := struct {
		User           *ent.User `json:"user"`
		FollowingCount int       `json:"following"`
		FollowersCount int       `json:"followers"`
	}{
		User:           user,
		FollowersCount: followersCount,
		FollowingCount: followingCount,
	}

	return c.JSON(http.StatusOK, res)
}

func (h UserHandler) GetFollowers(c echo.Context) error {
	// Get the user id from request parameter
	userID := c.Param("id")

	// Cache key specific to user and followers
	cacheKey := fmt.Sprintf("followers-%s", userID)

	// Retrieve followers from cache or database
	followers, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.User, error) {
		return h.db.User.Query().
			Where(user.ID(userID)).
			QueryFollowers().
			All(ctx)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch followers: " + err.Error()})
	}

	return c.JSON(http.StatusOK, followers)
}

func (h UserHandler) GetFollowing(c echo.Context) error {
	// Get the user id from request parameter
	userID := c.Param("id")

	// Cache key specific to user and followers
	cacheKey := fmt.Sprintf("following-%s", userID)

	// Retrieve following from cache or database
	following, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.User, error) {
		return h.db.User.Query().
			Where(user.ID(userID)).
			QueryFollowing().
			All(ctx)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch following: " + err.Error()})
	}

	return c.JSON(http.StatusOK, following)
}

func (h UserHandler) UpdateUserProfile(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Bind the request data to UserRequest
	var req requests.ProfileUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Update fields
	user, err := h.db.User.
		UpdateOneID(userID).
		SetName(req.Name).
		SetBio(req.Bio).
		SetURL(req.Url).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update user profile: " + err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) UpdateUserAccount(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Bind the request data to UserRequest
	var req requests.AcountUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request: " + err.Error()})
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Update fields
	user, err := h.db.User.
		UpdateOneID(userID).
		SetUsername(req.Username).
		SetEmail(req.Email).
		SetDob(req.DOB).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update user account: " + err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) DeleteUser(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Query post by user id
	post, err := h.db.Post.
		Query().
		Where(post.HasUserWith(user.IDEQ(userID))).
		Limit(1).
		All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// If there is more than one post by this user, prompt the user to delete the posts first
	if len(post) > 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "You must delete your posts before deleting the account"})
	}

	// Query user with user id
	user, err := h.db.User.
		Get(c.Request().Context(), userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// If the user has a profile image, delete it
	if user.ProfileImageURL != "" {
		err := storage.DeleteFile(c.Request().Context(), user.ProfileImageURL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user profile image"})
		}
	}

	// Delete the user
	err = h.db.User.
		DeleteOne(user).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete the user"})
	}

	return c.NoContent(http.StatusOK)
}

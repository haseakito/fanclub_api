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
	"github.com/hackgame-org/fanclub_api/pkg/storage"
	"github.com/labstack/echo/v4"
	muxgo "github.com/muxinc/mux-go"
)

type PostHandler struct {
	db    *ent.Client
	mux   *muxgo.APIClient
	cache *cache.Cache[[]*ent.Post]
}

func NewPostHandler(db *ent.Client, mux *muxgo.APIClient, rdb redis.Client) *PostHandler {
	return &PostHandler{
		db:  db,
		mux: mux,
		cache: cache.NewCache[[]*ent.Post](
			rdb,
			time.Minute,
		),
	}
}

func (h PostHandler) CreatePost(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

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
		SetUser(h.db.User.GetX(c.Request().Context(), userID)).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h PostHandler) UploadVideo(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

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
	defer src.Close()

	// Construct a unique key for the S3 object
	key := "videos/" + postID + "-" + file.Filename

	// Upload file to storage
	url, err := storage.UploadFile(c.Request().Context(), src, file.Header.Get("Content-Type"), key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload video: " + err.Error()})
	}

	// Create Mux asset for streaming
	asset, err := h.mux.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			{
				Url: url,
			},
		},
		PlaybackPolicy: []muxgo.PlaybackPolicy{
			muxgo.PUBLIC,
		},
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create the mux asset: " + err.Error()})
	}

	// Update fields
	_, err = h.db.Post.UpdateOneID(postID).
		SetVideoURL(url).
		SetMuxAssetID(asset.Data.Id).
		SetMuxPlaybackID(asset.Data.PlaybackIds[0].Id).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update the post: " + err.Error()})
	}

	return c.JSON(http.StatusOK, "Successfully uploaded the video")
}

func (h PostHandler) UploadThumnail(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

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
	defer src.Close()

	// Construct a unique key for the S3 object
	key := "thumnails/" + postID + "-" + file.Filename

	// Upload file to storage
	url, err := storage.UploadFile(c.Request().Context(), src, file.Header.Get("Content-Type"), key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload thumnail: " + err.Error()})
	}

	// Update fields
	_, err = h.db.Post.UpdateOneID(postID).
		SetThumbnailURL(url).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update the post: " + err.Error()})
	}

	return c.JSON(http.StatusOK, "Successfully uploaded the thumnail")
}

func (h PostHandler) GetPosts(c echo.Context) error {
	// Get user ID from query parameter
	userID := c.QueryParam("userId")

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
	var cacheKey string
	if userID != "" {
		cacheKey = fmt.Sprintf("posts-user-%s-limit-%d-offset-%d", userID, limit, offset)
	} else {
		cacheKey = fmt.Sprintf("posts-limit-%d-offset-%d", limit, offset)
	}

	// Query posts with limit and offset
	res, err := h.cache.GetOrSet(c.Request().Context(), cacheKey, func(ctx context.Context) ([]*ent.Post, error) {
		query := h.db.Post.Query().WithCategories().Limit(limit).Offset(offset)
		// Filter by user id if provided
		if userID != "" {
			query = query.Where(post.HasUserWith(user.IDEQ(userID)))
		}
		return query.All(ctx)
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch posts: " + err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h PostHandler) GetPostByID(c echo.Context) error {
	// Get the user id from context
	userID := c.Get("userID").(string)

	// Get post id from request
	postID := c.Param("id")

	// Query the post with post id
	postData, err := h.db.Post.
		Query().
		Where(post.ID(postID)).
		WithUser().
		WithCategories().
		WithSubscriptions().
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Query total likes for the post
	likeCount, err := h.db.Like.
		Query().
		Where(like.HasPostWith(post.ID(postID))).
		Count(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch a number of likes: " + err.Error()})
	}

	// Determine if the current user has liked the post
	userLiked, err := h.db.Like.
		Query().
		Where(
			like.HasUserWith(user.ID(userID)),
			like.HasPostWith(post.ID(postID)),
		).
		Exist(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch the like" + err.Error()})
	}

	// Create the response object
	res := struct {
		Post      *ent.Post `json:"post"`
		User      *ent.User `json:"user"`
		LikeCount int       `json:"likes"`
		UserLiked bool      `json:"userLiked"`
	}{
		Post:      postData,
		User:      postData.Edges.User,
		LikeCount: likeCount,
		UserLiked: userLiked,
	}

	return c.JSON(http.StatusOK, res)
}

func (h PostHandler) UpdatePost(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

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
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Update fields with data provided
	res, err := tx.Post.
		UpdateOneID(postID).
		SetTitle(req.Title).
		SetDescription(*req.Description).
		SetPrice(*req.Price).
		SetIsFeatured(req.IsFeatured).
		SetStatus(req.Status).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update fields: " + err.Error()})
	}

	// Iterate through categories to associate the post with categories
	for _, categoryID := range req.CategoryIDs {
		// Query the category with category id
		category, err := tx.Category.Get(c.Request().Context(), categoryID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch category: " + err.Error()})
		}

		// Update the category field for the post
		if _, err = res.Update().
			AddCategories(category).
			Save(c.Request().Context()); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed add category to post: " + err.Error()})
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

	// Query a post with the post id
	post, err := h.db.Post.Get(c.Request().Context(), postID)
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// If the post has a video, delete it
	if post.VideoURL != "" {
		// Delete video from storage
		err = storage.DeleteFile(c.Request().Context(), post.VideoURL)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete video: "+err.Error())
		}

		// Delete mux asset associated with the video
		err = h.mux.AssetsApi.DeleteAsset(post.MuxAssetID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to delete mux asset: "+err.Error())
		}
	}

	// If the post has a thumnail, delete it
	if post.ThumbnailURL != "" {
		storage.DeleteFile(c.Request().Context(), post.ThumbnailURL)
	}

	// Delete the post
	err = h.db.Post.
		DeleteOneID(postID).
		Exec(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Successfully delete the post")
}

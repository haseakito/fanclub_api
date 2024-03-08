package handlers

import (
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/asset"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	db  *ent.Client
	cld *cloudinary.Cloudinary
}

func NewPostHandler(db *ent.Client, cld *cloudinary.Cloudinary) *PostHandler {
	return &PostHandler{
		db:  db,
		cld: cld,
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
		SetUser(h.db.User.GetX(c.Request().Context(), req.UserID)).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h PostHandler) UploadFiles(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	// Get files from request
	files := form.File["images"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, "No files uploaded")
	}

	// Start a new transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Iterate over images and upload them to cloudinary
	for _, file := range files {
		// Validate request file
		if err := requests.ValidateFile(file); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		// Upload an asset to cloudinary
		resp, err := h.cld.Upload.Upload(
			c.Request().Context(),
			src,
			uploader.UploadParams{},
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// Insert a new asset
		asset, err := tx.Asset.
			Create().
			SetPublicID(resp.PublicID).
			SetURL(resp.SecureURL).
			SetResourceType(resp.ResourceType).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// Update asset field
		_, err = tx.Post.
			UpdateOneID(postUUID).
			AddAssets(asset).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Successfully uploaded the images")
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
		WithAssets().
		WithCategories().
		Limit(limit).
		Offset(offset).
		All(c.Request().Context())
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
		WithAssets().
		Only(c.Request().Context())
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
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Update fields with data provided
	res, err := tx.Post.
		UpdateOneID(postUUID).
		SetTitle(req.Title).		
		SetDescription(*req.Description).
		SetPrice(*req.Price).
		SetIsFeatured(req.IsFeatured).
		SetStatus(req.Status).
		Save(c.Request().Context())
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
		category, err := tx.Category.Get(c.Request().Context(), categoryUUID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		// Update the category field for the post
		if _, err = res.Update().
			AddCategories(category).
			Save(c.Request().Context()); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func (h PostHandler) DeleteFile(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")
	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Get asset id from request
	assetID := c.Param("asset_id")
	// Parse post ID string into UUID
	assetUUID, err := uuid.Parse(assetID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Query the asset with post id and asset id
	asset, err := h.db.Asset.
		Query().
		Where(asset.HasPostWith(post.ID(postUUID))).
		Where(asset.ID(assetUUID)).
		Only(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Delete the object from cloudinary
	_, err = h.cld.Upload.Destroy(c.Request().Context(), uploader.DestroyParams{
		PublicID:     asset.PublicID,
		ResourceType: asset.ResourceType,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Delete the asset with asset id
	err = h.db.Asset.
		DeleteOneID(assetUUID).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Successfully deleted the image")
}

func (h PostHandler) DeletePost(c echo.Context) error {
	// Get post id from request
	postID := c.Param("id")

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Start a new transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Query the billboard with billboard id
	post, err := tx.Post.
		Query().
		Where(post.ID(postUUID)).
		WithAssets().
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Iterate over assets and delete the corresponding object from cloudinary
	if len(post.Edges.Assets) > 0 {
		for _, asset := range post.Edges.Assets {
			// Delete the object from cloudinary
			_, err = h.cld.Upload.Destroy(c.Request().Context(), uploader.DestroyParams{
				PublicID:     asset.PublicID,
				ResourceType: asset.ResourceType,
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
	}

	// Delete the post with post id
	if err := h.db.Post.
		DeleteOneID(postUUID).
		Exec(c.Request().Context()); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete post")
	}

	return c.JSON(http.StatusOK, "Successfully delete the post")
}

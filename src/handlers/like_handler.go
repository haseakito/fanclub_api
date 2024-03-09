package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/like"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/ent/user"
	"github.com/hackgame-org/fanclub_api/requests"
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

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(req.PostID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Query the like with user id and post id
	_, err = h.db.Like.
		Query().
		Where(like.HasPostWith(post.ID(postUUID))).
		Where(like.HasUserWith(user.ID(req.UserID))).
		First(c.Request().Context())
	if err != nil {
		// If there is no like associated with user id and post id, then proceed to make a new one
		if ent.IsNotFound(err) {
			// Query the post with post id
			postData, err := h.db.Post.
				Query().
				Where(post.ID(postUUID)).
				WithUser().
				Only(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			// If the post is owned by the user, then throw an error
			if postData.Edges.User.ID == req.UserID {
				return c.String(http.StatusUnauthorized, "Unauthorized to perform this operation")
			}

			// Insert a new like with user id and post id
			_, err = h.db.Like.
				Create().
				SetUser(h.db.User.GetX(c.Request().Context(), req.UserID)).
				SetPost(postData).
				Save(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusCreated, "Successfully liked the post")
		}

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "Already liked the post")
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

	// Parse post ID string into UUID
	postUUID, err := uuid.Parse(req.PostID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Delete the like with post id and user id
	_, err = h.db.Like.
		Delete().
		Where(like.HasPostWith(post.ID(postUUID))).
		Where(like.HasUserWith(user.ID(req.UserID))).
		Exec(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "Successfully deleted the like")
}

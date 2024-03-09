package ent_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/enttest"
	"github.com/hackgame-org/fanclub_api/ent/like"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/ent/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLike(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a know ID for user
	userID := uuid.NewString()
	// Generate a known UUID for the post
	postUUID := uuid.New()

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID(userID).
		SetName("test user").
		SetEmail("example@example.com").
		Save(ctx)
	require.NoError(t, err)

	// Insert a new post
	_, err = tx.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").
		SetUser(user).
		Save(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Generate a known UUID for the like
	likeUUID := uuid.New()

	// Insert a new like with user id and post id
	_, err = client.Like.
		Create().
		SetID(likeUUID).
		SetPost(client.Post.GetX(ctx, postUUID)).
		SetUser(client.User.GetX(ctx, userID)).
		Save(ctx)
	require.NoError(t, err)

	// Query the like with like id
	like, err := client.Like.
		Query().
		WithUser().
		WithPost().
		Only(ctx)
	require.NoError(t, err)

	// Assert that the returned like is not nil
	assert.NotNil(t, like)

	// Assert that the returned like has the expected user and post relation
	assert.Equal(t, userID, like.Edges.User.ID)
	assert.Equal(t, postUUID, like.Edges.Post.ID)
}

func TestDeleteLike(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a know ID for user
	userID := strconv.Itoa(rand.Intn(1000))
	// Generate a known UUID for the post
	postUUID := uuid.New()
	// Generate a known UUID for the like
	likeUUID := uuid.New()

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Insert a new user
	usr, err := tx.User.
		Create().
		SetID(userID).
		SetName("test user").
		SetEmail("example@example.com").
		Save(ctx)
	require.NoError(t, err)

	// Insert a new post
	_, err = tx.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").
		SetUser(usr).
		Save(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Insert a new like
	_, err = client.Like.
		Create().
		SetID(likeUUID).
		SetPost(client.Post.GetX(ctx, postUUID)).
		SetUser(client.User.GetX(ctx, userID)).
		Save(ctx)
	require.NoError(t, err)

	// Delete the test data with post id
	_, err = client.Like.
		Delete().
		Where(like.HasPostWith(post.ID(postUUID))).
		Where(like.HasUserWith(user.ID(userID))).
		Exec(ctx)
	require.NoError(t, err)
}

package ent_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/enttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateFollow(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate know IDs for user
	userID1 := uuid.NewString()
	userID2 := uuid.NewString()

	// Insert a new user
	_, err = client.User.
		CreateBulk(
			client.User.Create().
				SetID(userID1).
				SetName("test user 1").
				SetEmail("example1@example.com"),
			client.User.Create().
				SetID(userID2).
				SetName("test user 2").
				SetEmail("example2@example.com"),
		).
		Save(ctx)
	require.NoError(t, err)

	// Insert a new follow with follower id and following user id
	user, err := client.User.
		UpdateOneID(userID1).
		AddFollowers(client.User.GetX(ctx, userID2)).
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned user is not nil
	assert.NotNil(t, user)

	// Query follower
	follower := user.QueryFollowers().OnlyX(ctx)

	// Assert that the returned follower matches the expected user id
	assert.Equal(t, userID2, follower.ID)
}

func TestDeleteFollow(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate know IDs for user
	userID1 := uuid.NewString()
	userID2 := uuid.NewString()

	// Insert a new user
	_, err = client.User.
		CreateBulk(
			client.User.Create().
				SetID(userID1).
				SetName("test user 1").
				SetEmail("example1@example.com"),
			client.User.Create().
				SetID(userID2).
				SetName("test user 2").
				SetEmail("example2@example.com"),
		).
		Save(ctx)
	require.NoError(t, err)

	// Insert a new follow with follower id and following user id
	user, err := client.User.
		UpdateOneID(userID1).
		AddFollowers(client.User.GetX(ctx, userID2)).
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned user is not nil
	assert.NotNil(t, user)

	// Unfollow
	_, err = client.User.
		UpdateOneID(userID1).
		RemoveFollowers(client.User.GetX(ctx, userID2)).
		Save(ctx)
	require.NoError(t, err)

	// Query followers
	followers := user.QueryFollowers().AllX(ctx)

	// Assert that the returned followers are empty
	assert.Empty(t, followers)
}

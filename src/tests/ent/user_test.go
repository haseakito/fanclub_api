package ent_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/api/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate known IDs for users
	userID1 := uuid.NewString()
	userID2 := uuid.NewString()

	// Insert test data
	_, err = client.User.
		CreateBulk(
			client.User.
				Create().
				SetID(userID1).
				SetName("test user 1").
				SetEmail("example1@example.com").
				SetPassword("test user password"),
			client.User.
				Create().
				SetID(userID2).
				SetName("test user 2").
				SetEmail("example2@example.com").
				SetPassword("test user password"),
		).
		Save(ctx)
	require.NoError(t, err)

	// Get all users
	users, err := client.User.Query().All(ctx)
	require.NoError(t, err)

	// Assert that the returned user id matches the expected id
	assert.Equal(t, userID1, users[0].ID)
	assert.Equal(t, userID2, users[1].ID)
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a know ID for user
	userID := strconv.Itoa(rand.Intn(1000))

	// Insert test data
	_, err = client.User.
		Create().
		SetID(userID).
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Query the user with user id
	user, err := client.User.Get(ctx, userID)
	require.NoError(t, err)

	// Assert that the returned user is not nil
	assert.NotNil(t, user)

	// Assert that the returned user name matches the expected name
	assert.Equal(t, "test user", user.Name)
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a know ID for user
	userID := strconv.Itoa(rand.Intn(1000))

	// Insert test data
	_, err = client.User.
		Create().
		SetID(userID).
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Update fields
	user, err := client.User.
		UpdateOneID(userID).
		SetName("modified test user").
		SetBio("test bio").
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned user is not nil
	assert.NotNil(t, user)

	// Assert that the returned user name matches the expected name
	assert.Equal(t, "modified test user", user.Name)
	// Assert that the returned user bio matches the expected bio
	assert.Equal(t, "test bio", user.Bio)
}

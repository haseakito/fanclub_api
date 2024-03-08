package ent_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate known UUIDs for categories
	categoryUUID1 := uuid.New()
	categoryUUID2 := uuid.New()

	// Insert test data
	_, err = client.Category.
		CreateBulk(
			client.Category.Create().SetID(categoryUUID1).SetName("test category 1"),
			client.Category.Create().SetID(categoryUUID2).SetName("test category 2"),
		).
		Save(ctx)
	require.NoError(t, err)

	// Get all categories
	categories, err := client.Category.Query().All(ctx)
	require.NoError(t, err)

	// Assert that the returned categories is not nil
	require.NotNil(t, categories)

	// Assert that the returned category id matches the expected id
	assert.Equal(t, categoryUUID1, categories[0].ID)
	assert.Equal(t, categoryUUID2, categories[1].ID)
}

func TestGetPostsByCategoryID(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Generate known UUID for categories
	categoryUUID := uuid.New()

	// Insert test data
	_, err = tx.Category.
		Create().
		SetID(categoryUUID).
		SetName("test category").
		Save(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the post
	postUUID := uuid.New()

	// Insert test data
	_, err = tx.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").		
		AddCategories(tx.Category.GetX(ctx, categoryUUID)).
		Save(ctx)
	require.NoError(t, err)

	// Query the posts with category id
	posts, err := tx.Category.
		QueryPosts(tx.Category.GetX(ctx, categoryUUID)).
		All(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Assert that the fetched posts exist
	assert.NotNil(t, posts)

	// Assert that the returned posts matches the expected post
	assert.Equal(t, postUUID, posts[0].ID)
}

func TestDeleteCategory(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the category
	categoryUUID := uuid.New()

	// Insert test data
	_, err = client.Category.
		Create().
		SetID(categoryUUID).
		SetName("test category").
		Save(context.Background())
	require.NoError(t, err)

	// Delete the test data with post id
	err = client.Category.
		DeleteOneID(categoryUUID).
		Exec(ctx)
	require.NoError(t, err)
}

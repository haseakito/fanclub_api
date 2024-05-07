package ent_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/api/ent/enttest"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
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

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID("test-id").
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Insert test data
	post, err := tx.Post.
		Create().
		SetTitle("test title").
		SetUser(user).
		Save(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Assert that the returned post is not nil
	assert.NotNil(t, post)

	// Assert that the returned post title matches the expected title
	assert.Equal(t, "test title", post.Title)
}

func TestCreatePostWithCategory(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the category
	categoryUUID := uuid.NewString()

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID("test-id").
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Insert test category
	_, err = tx.Category.
		Create().
		SetID(categoryUUID).
		SetName("test category").
		Save(ctx)
	require.NoError(t, err)

	// Insert test post
	newPost, err := tx.Post.
		Create().
		SetTitle("test post title").
		SetUser(user).
		SetDescription("").
		SetPrice(0).
		AddCategories(tx.Category.GetX(ctx, categoryUUID)).
		Save(ctx)
	require.NoError(t, err)

	// Get the post together with categories with post id
	postWithCategories, err := tx.Post.
		Query().
		Where(post.ID(newPost.ID)).
		WithCategories().
		Only(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Assert that the fetched post contains the expected category
	assert.NotNil(t, postWithCategories)

	// Assert that the returned post category matches the expected category
	assert.Equal(t, categoryUUID, postWithCategories.Edges.Categories[0].ID)
}

func TestCreatePostWithSubscription(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the subscription
	subscriptionUUID := uuid.NewString()

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID("test-id").
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Insert test subscription
	_, err = tx.Subscription.
		Create().
		SetID(subscriptionUUID).
		SetName("test subscription name").
		SetUser(user).
		SetDescription("").
		SetPrice(0).
		SetTrialPeriodDays(0).
		Save(ctx)
	require.NoError(t, err)

	// Insert test post
	newPost, err := tx.Post.
		Create().
		SetTitle("test post title").
		SetUser(user).
		SetDescription("").
		SetPrice(0).
		AddSubscriptions(tx.Subscription.GetX(ctx, subscriptionUUID)).
		Save(ctx)
	require.NoError(t, err)

	// Get the post together with subscriptions with post id
	postWithSubscriptions, err := tx.Post.
		Query().
		Where(post.ID(newPost.ID)).
		WithSubscriptions().
		Only(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Assert that the fetched post contains the expected subscription
	assert.NotNil(t, postWithSubscriptions)

	// Assert that the returned post subscription matches the expected subscription
	assert.Equal(t, subscriptionUUID, postWithSubscriptions.Edges.Subscriptions[0].ID)
}

func TestGetPosts(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate known UUIDs for posts
	postUUID1 := uuid.NewString()
	postUUID2 := uuid.NewString()

	// Insert test data
	_, err = client.Post.
		CreateBulk(
			client.Post.
				Create().
				SetID(postUUID1).
				SetTitle("test title 1").
				SetDescription("").
				SetPrice(0),
			client.Post.
				Create().
				SetID(postUUID2).
				SetTitle("test title 2").
				SetDescription("").
				SetPrice(0),
		).
		Save(ctx)
	require.NoError(t, err)

	// Get all posts
	posts, err := client.Post.Query().All(ctx)
	require.NoError(t, err)

	// Assert that the returned posts is not nil
	require.NotNil(t, posts)

	// Assert that the returned post id matches the expected id
	assert.Equal(t, postUUID1, posts[0].ID)
	assert.Equal(t, postUUID2, posts[1].ID)
}

func TestGetPost(t *testing.T) {
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

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID("test-id").
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the post
	postUUID := uuid.NewString()

	// Insert test data
	_, err = tx.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").
		SetUser(user).
		SetDescription("").
		SetPrice(0).
		SetIsFeatured(false).
		SetStatus(false).
		Save(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Get the post with post id
	post, err := client.Post.Get(ctx, postUUID)
	require.NoError(t, err)

	// Assert that the returned post is not nil
	require.NotNil(t, post)

	// Assert that the returned post title matches the expected title
	assert.Equal(t, "test title", post.Title)
}

func TestUpdatePost(t *testing.T) {
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

	// Insert a new user
	user, err := tx.User.
		Create().
		SetID("test-id").
		SetName("test user").
		SetEmail("example@example.com").
		SetPassword("test user password").
		Save(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the post
	postUUID := uuid.NewString()

	// Insert test data
	_, err = tx.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").
		SetUser(user).
		SetDescription("test description").
		SetPrice(0).
		SetIsFeatured(false).
		SetStatus(false).
		Save(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Update the test data with post id
	post, err := client.Post.
		UpdateOneID(postUUID).
		SetTitle("modified test title").
		SetDescription("modified test description").
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned post is not nil
	require.NotNil(t, post)

	// Assert that the returned post title matches the expected title
	assert.Equal(t, "modified test title", post.Title)
	// Assert that the returned post description matches the expected description
	assert.Equal(t, "modified test description", post.Description)
}

func TestDeletePost(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the post
	postUUID := uuid.NewString()

	// Insert test data
	_, err = client.Post.
		Create().
		SetID(postUUID).
		SetTitle("test title").
		SetDescription("test description").
		SetPrice(0).
		SetIsFeatured(false).
		SetStatus(false).
		Save(context.Background())
	require.NoError(t, err)

	// Delete the test data with post id
	err = client.Post.
		DeleteOneID(postUUID).
		Exec(ctx)
	require.NoError(t, err)
}

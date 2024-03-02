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

func TestCreateSubscription(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Insert test data
	sub, err := client.Subscription.
		Create().
		SetName("test subscription").
		SetUserID("test-user-id").
		SetDescription("").
		SetPrice(0).
		SetTrialPeriodDays(0).
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned subscription is not nil
	assert.NotNil(t, sub)

	// Assert that the returned subscription title matches the expected title
	assert.Equal(t, "test subscription", sub.Name)
}

func TestGetSubscriptions(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)
}

func TestGetSubscription(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the subscription
	subUUID := uuid.New()

	// Insert test data
	_, err = client.Subscription.
		Create().
		SetID(subUUID).
		SetName("test subscription").
		SetUserID("test-user-id").
		SetDescription("").
		SetPrice(0).
		SetTrialPeriodDays(0).
		Save(ctx)
	require.NoError(t, err)

	// Get the subscription with subscription id
	sub, err := client.Subscription.Get(ctx, subUUID)
	require.NoError(t, err)

	// Assert that the returned subscription is not nil
	require.NotNil(t, sub)

	// Assert that the returned subscription id matches the expected id
	assert.Equal(t, subUUID, sub.ID)
	// Assert that the returned subscription title matches the expected title
	assert.Equal(t, "test subscription", sub.Name)
}

func TestUpdateSubscription(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the subscription
	subUUID := uuid.New()

	// Insert test data
	_, err = client.Subscription.
		Create().
		SetID(subUUID).
		SetName("test subscription").
		SetUserID("test-user-id").
		SetDescription("").
		SetPrice(0).
		SetTrialPeriodDays(0).
		Save(ctx)
	require.NoError(t, err)

	// Update the test data with subscription id
	sub, err := client.Subscription.
		UpdateOneID(subUUID).
		SetName("modified test subscription").
		SetDescription("modified test description").
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned subscription is not nil
	require.NotNil(t, sub)

	// Assert that the returned subscription title matches the expected title
	assert.Equal(t, "modified test subscription", sub.Name)
	// Assert that the returned subscription description matches the expected description
	assert.Equal(t, "modified test description", *sub.Description)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for the subscription
	subUUID := uuid.New()

	// Insert test data
	_, err = client.Subscription.
		Create().
		SetID(subUUID).
		SetName("test subscription").
		SetUserID("test-user-id").
		SetDescription("").
		SetPrice(0).
		SetTrialPeriodDays(0).
		Save(ctx)
	require.NoError(t, err)

	// Delete the test data with subscription id
	err = client.Subscription.
		DeleteOneID(subUUID).
		Exec(ctx)
	require.NoError(t, err)
}

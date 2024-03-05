package ent_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/billboard"
	"github.com/hackgame-org/fanclub_api/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBillboard(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Insert a new billboard
	bill, err := client.Billboard.
		Create().
		SetTitle("test title").
		SetDescription("test description").
		Save(ctx)
	require.NoError(t, err)

	// Assert that the returned billboard is not nil
	assert.NotNil(t, bill)

	// Assert that the returned billboard title matches the expected title
	assert.Equal(t, "test title", bill.Title)
}

func TestUploadFile(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for billboard
	billUUID := uuid.New()

	// Start a new transaction
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx.Rollback()

	// Insert a new billboard
	_, err = tx.Billboard.
		Create().
		SetID(billUUID).
		SetTitle("test upload").
		Save(ctx)
	require.NoError(t, err)

	// Insert a new asset
	asset, err := tx.Asset.
		Create().
		SetPublicID("test public id").		
		SetResourceType("image").
		SetURL("https://res.cloudinary.com/demo/image/upload/v1591095352/hl22acprlomnycgiudor.jpg").
		Save(ctx)
	require.NoError(t, err)

	// Update asset field
	_, err = tx.Billboard.
		UpdateOneID(billUUID).
		SetAsset(asset).
		Save(ctx)
	require.NoError(t, err)

	// Query billboard
	bill, err := tx.Billboard.Query().WithAsset().Only(ctx)
	require.NoError(t, err)

	// Commit the transaction
	err = tx.Commit()
	require.NoError(t, err)

	// Assert that the returned billboard is not nil
	assert.NotNil(t, bill)

	// Assert that the returned billboard asset edge is not nil
	assert.NotNil(t, bill.Edges.Asset)
}

func TestGetBillboards(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate known UUIDs for billboards
	billUUID1 := uuid.New()
	billUUID2 := uuid.New()

	// Insert test data
	_, err = client.Billboard.
		CreateBulk(
			client.Billboard.
				Create().
				SetID(billUUID1).
				SetTitle("test title 1").
				SetDescription("test description"),
			client.Billboard.
				Create().
				SetID(billUUID2).
				SetTitle("test title 2").
				SetDescription("test description"),
		).
		Save(ctx)
	require.NoError(t, err)

	// Get all billboards
	bills, err := client.Billboard.Query().All(ctx)
	require.NoError(t, err)

	// Assert that the returned billboards are not nil
	assert.NotNil(t, bills)

	// Assert that the returned billboard id matches the expected id
	assert.Equal(t, billUUID1, bills[0].ID)
	assert.Equal(t, billUUID2, bills[1].ID)
}

func TestGetBillboard(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for billboard
	billUUID := uuid.New()

	// Insert test data
	_, err = client.Billboard.
		Create().
		SetID(billUUID).
		SetTitle("").
		SetDescription("").
		Save(ctx)
	require.NoError(t, err)

	// Get all billboards
	bill, err := client.Billboard.
		Query().
		Where(billboard.ID(billUUID)).
		Only(ctx)
	require.NoError(t, err)

	// Assert that the returned billboard is not nil
	assert.NotNil(t, bill)

	// Assert that the returned billboard id matches the expected id
	assert.Equal(t, billUUID, bill.ID)
}

func TestDeleteBillboard(t *testing.T) {
	ctx := context.Background()

	// Open a test SQLite database
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create the schema
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// Generate a known UUID for billboard
	billUUID := uuid.New()

	// Insert test data
	_, err = client.Billboard.
		Create().
		SetID(billUUID).
		SetTitle("").
		SetDescription("").
		Save(ctx)
	require.NoError(t, err)

	// Delete the test data with billboard id
	err = client.Billboard.
		DeleteOneID(billUUID).
		Exec(ctx)
	require.NoError(t, err)
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/category"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/ent/schema"
	"github.com/hackgame-org/fanclub_api/ent/subscription"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[2].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryFields[3].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryFields[0].Descriptor()
	// category.DefaultID holds the default value on creation for the id field.
	category.DefaultID = categoryDescID.Default.(func() uuid.UUID)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescIsFeatured is the schema descriptor for is_featured field.
	postDescIsFeatured := postFields[5].Descriptor()
	// post.DefaultIsFeatured holds the default value on creation for the is_featured field.
	post.DefaultIsFeatured = postDescIsFeatured.Default.(bool)
	// postDescStatus is the schema descriptor for status field.
	postDescStatus := postFields[6].Descriptor()
	// post.DefaultStatus holds the default value on creation for the status field.
	post.DefaultStatus = postDescStatus.Default.(bool)
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[7].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[8].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.DefaultID holds the default value on creation for the id field.
	post.DefaultID = postDescID.Default.(func() uuid.UUID)
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescIsArchived is the schema descriptor for is_archived field.
	subscriptionDescIsArchived := subscriptionFields[6].Descriptor()
	// subscription.DefaultIsArchived holds the default value on creation for the is_archived field.
	subscription.DefaultIsArchived = subscriptionDescIsArchived.Default.(bool)
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionFields[7].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() time.Time)
	// subscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	subscriptionDescUpdatedAt := subscriptionFields[8].Descriptor()
	// subscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	subscription.DefaultUpdatedAt = subscriptionDescUpdatedAt.Default.(func() time.Time)
	// subscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	subscription.UpdateDefaultUpdatedAt = subscriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// subscriptionDescID is the schema descriptor for id field.
	subscriptionDescID := subscriptionFields[0].Descriptor()
	// subscription.DefaultID holds the default value on creation for the id field.
	subscription.DefaultID = subscriptionDescID.Default.(func() uuid.UUID)
}
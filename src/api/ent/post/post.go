// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldThumbnailURL holds the string denoting the thumbnail_url field in the database.
	FieldThumbnailURL = "thumbnail_url"
	// FieldVideoURL holds the string denoting the video_url field in the database.
	FieldVideoURL = "video_url"
	// FieldMuxAssetID holds the string denoting the mux_asset_id field in the database.
	FieldMuxAssetID = "mux_asset_id"
	// FieldMuxPlaybackID holds the string denoting the mux_playback_id field in the database.
	FieldMuxPlaybackID = "mux_playback_id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldIsFeatured holds the string denoting the is_featured field in the database.
	FieldIsFeatured = "is_featured"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSubscriptions holds the string denoting the subscriptions edge name in mutations.
	EdgeSubscriptions = "subscriptions"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "posts"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_posts"
	// SubscriptionsTable is the table that holds the subscriptions relation/edge. The primary key declared below.
	SubscriptionsTable = "subscription_posts"
	// SubscriptionsInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionsInverseTable = "subscriptions"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "post_likes"
	// CategoriesTable is the table that holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "post_categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldThumbnailURL,
	FieldVideoURL,
	FieldMuxAssetID,
	FieldMuxPlaybackID,
	FieldPrice,
	FieldIsFeatured,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_posts",
}

var (
	// SubscriptionsPrimaryKey and SubscriptionsColumn2 are the table columns denoting the
	// primary key for the subscriptions relation (M2M).
	SubscriptionsPrimaryKey = []string{"subscription_id", "post_id"}
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"post_id", "category_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsFeatured holds the default value on creation for the "is_featured" field.
	DefaultIsFeatured bool
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Post queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByThumbnailURL orders the results by the thumbnail_url field.
func ByThumbnailURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailURL, opts...).ToFunc()
}

// ByVideoURL orders the results by the video_url field.
func ByVideoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoURL, opts...).ToFunc()
}

// ByMuxAssetID orders the results by the mux_asset_id field.
func ByMuxAssetID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMuxAssetID, opts...).ToFunc()
}

// ByMuxPlaybackID orders the results by the mux_playback_id field.
func ByMuxPlaybackID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMuxPlaybackID, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByIsFeatured orders the results by the is_featured field.
func ByIsFeatured(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFeatured, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// BySubscriptionsCount orders the results by subscriptions count.
func BySubscriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubscriptionsStep(), opts...)
	}
}

// BySubscriptions orders the results by subscriptions terms.
func BySubscriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubscriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikesCount orders the results by likes count.
func ByLikesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikesStep(), opts...)
	}
}

// ByLikes orders the results by likes terms.
func ByLikes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoriesCount orders the results by categories count.
func ByCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoriesStep(), opts...)
	}
}

// ByCategories orders the results by categories terms.
func ByCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newSubscriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubscriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SubscriptionsTable, SubscriptionsPrimaryKey...),
	)
}
func newLikesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikesTable, LikesColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CategoriesTable, CategoriesPrimaryKey...),
	)
}
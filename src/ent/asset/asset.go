// Code generated by ent, DO NOT EDIT.

package asset

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the asset type in the database.
	Label = "asset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPublicID holds the string denoting the public_id field in the database.
	FieldPublicID = "public_id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldResourceType holds the string denoting the resource_type field in the database.
	FieldResourceType = "resource_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBillboard holds the string denoting the billboard edge name in mutations.
	EdgeBillboard = "billboard"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// Table holds the table name of the asset in the database.
	Table = "assets"
	// BillboardTable is the table that holds the billboard relation/edge.
	BillboardTable = "assets"
	// BillboardInverseTable is the table name for the Billboard entity.
	// It exists in this package in order to avoid circular dependency with the "billboard" package.
	BillboardInverseTable = "billboards"
	// BillboardColumn is the table column denoting the billboard relation/edge.
	BillboardColumn = "billboard_asset"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "assets"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_assets"
)

// Columns holds all SQL columns for asset fields.
var Columns = []string{
	FieldID,
	FieldPublicID,
	FieldURL,
	FieldResourceType,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "assets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"billboard_asset",
	"post_assets",
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Asset queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPublicID orders the results by the public_id field.
func ByPublicID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByResourceType orders the results by the resource_type field.
func ByResourceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResourceType, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBillboardField orders the results by billboard field.
func ByBillboardField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBillboardStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostField orders the results by post field.
func ByPostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), sql.OrderByField(field, opts...))
	}
}
func newBillboardStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BillboardInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, BillboardTable, BillboardColumn),
	)
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
	)
}
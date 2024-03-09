// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent/post"
	"github.com/hackgame-org/fanclub_api/ent/user"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// IsFeatured holds the value of the "is_featured" field.
	IsFeatured *bool `json:"is_featured,omitempty"`
	// Status holds the value of the "status" field.
	Status *bool `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges        PostEdges `json:"edges"`
	user_posts   *string
	selectValues sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Assets holds the value of the assets edge.
	Assets []*Asset `json:"assets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[1] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[2] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[3] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) AssetsOrErr() ([]*Asset, error) {
	if e.loadedTypes[4] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldIsFeatured, post.FieldStatus:
			values[i] = new(sql.NullBool)
		case post.FieldPrice:
			values[i] = new(sql.NullInt64)
		case post.FieldTitle, post.FieldDescription:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case post.FieldID:
			values[i] = new(uuid.UUID)
		case post.ForeignKeys[0]: // user_posts
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case post.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				po.Description = value.String
			}
		case post.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				po.Price = int(value.Int64)
			}
		case post.FieldIsFeatured:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_featured", values[i])
			} else if value.Valid {
				po.IsFeatured = new(bool)
				*po.IsFeatured = value.Bool
			}
		case post.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = new(bool)
				*po.Status = value.Bool
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_posts", values[i])
			} else if value.Valid {
				po.user_posts = new(string)
				*po.user_posts = value.String
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Post entity.
func (po *Post) QueryUser() *UserQuery {
	return NewPostClient(po.config).QueryUser(po)
}

// QuerySubscriptions queries the "subscriptions" edge of the Post entity.
func (po *Post) QuerySubscriptions() *SubscriptionQuery {
	return NewPostClient(po.config).QuerySubscriptions(po)
}

// QueryLikes queries the "likes" edge of the Post entity.
func (po *Post) QueryLikes() *LikeQuery {
	return NewPostClient(po.config).QueryLikes(po)
}

// QueryCategories queries the "categories" edge of the Post entity.
func (po *Post) QueryCategories() *CategoryQuery {
	return NewPostClient(po.config).QueryCategories(po)
}

// QueryAssets queries the "assets" edge of the Post entity.
func (po *Post) QueryAssets() *AssetQuery {
	return NewPostClient(po.config).QueryAssets(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(po.Description)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", po.Price))
	builder.WriteString(", ")
	if v := po.IsFeatured; v != nil {
		builder.WriteString("is_featured=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

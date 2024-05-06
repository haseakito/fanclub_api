// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hackgame-org/fanclub_api/api/ent/like"
	"github.com/hackgame-org/fanclub_api/api/ent/post"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
)

// Like is the model entity for the Like schema.
type Like struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LikeQuery when eager-loading is set.
	Edges        LikeEdges `json:"edges"`
	post_likes   *string
	user_likes   *string
	selectValues sql.SelectValues
}

// LikeEdges holds the relations/edges for other nodes in the graph.
type LikeEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) PostOrErr() (*Post, error) {
	if e.Post != nil {
		return e.Post, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: post.Label}
	}
	return nil, &NotLoadedError{edge: "post"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Like) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case like.FieldID:
			values[i] = new(sql.NullInt64)
		case like.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case like.ForeignKeys[0]: // post_likes
			values[i] = new(sql.NullString)
		case like.ForeignKeys[1]: // user_likes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Like fields.
func (l *Like) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case like.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case like.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case like.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_likes", values[i])
			} else if value.Valid {
				l.post_likes = new(string)
				*l.post_likes = value.String
			}
		case like.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_likes", values[i])
			} else if value.Valid {
				l.user_likes = new(string)
				*l.user_likes = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Like.
// This includes values selected through modifiers, order, etc.
func (l *Like) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryPost queries the "post" edge of the Like entity.
func (l *Like) QueryPost() *PostQuery {
	return NewLikeClient(l.config).QueryPost(l)
}

// QueryUser queries the "user" edge of the Like entity.
func (l *Like) QueryUser() *UserQuery {
	return NewLikeClient(l.config).QueryUser(l)
}

// Update returns a builder for updating this Like.
// Note that you need to call Like.Unwrap() before calling this method if this Like
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Like) Update() *LikeUpdateOne {
	return NewLikeClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Like entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Like) Unwrap() *Like {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Like is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Like) String() string {
	var builder strings.Builder
	builder.WriteString("Like(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Likes is a parsable slice of Like.
type Likes []*Like
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/ent/verificationtoken"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// ProfileImageURL holds the value of the "profile_image_url" field.
	ProfileImageURL string `json:"profile_image_url,omitempty"`
	// StripeCustomerID holds the value of the "stripe_customer_id" field.
	StripeCustomerID string `json:"stripe_customer_id,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified *bool `json:"email_verified,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// Dob holds the value of the "dob" field.
	Dob *string `json:"dob,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// VerificationToken holds the value of the verification_token edge.
	VerificationToken *VerificationToken `json:"verification_token,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// VerificationTokenOrErr returns the VerificationToken value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) VerificationTokenOrErr() (*VerificationToken, error) {
	if e.VerificationToken != nil {
		return e.VerificationToken, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: verificationtoken.Label}
	}
	return nil, &NotLoadedError{edge: "verification_token"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[1] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[2] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[3] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailVerified:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldName, user.FieldUsername, user.FieldProfileImageURL, user.FieldStripeCustomerID, user.FieldPassword, user.FieldURL, user.FieldEmail, user.FieldBio, user.FieldDob:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldProfileImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image_url", values[i])
			} else if value.Valid {
				u.ProfileImageURL = value.String
			}
		case user.FieldStripeCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stripe_customer_id", values[i])
			} else if value.Valid {
				u.StripeCustomerID = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				u.EmailVerified = new(bool)
				*u.EmailVerified = value.Bool
			}
		case user.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				u.Bio = value.String
			}
		case user.FieldDob:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dob", values[i])
			} else if value.Valid {
				u.Dob = new(string)
				*u.Dob = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryVerificationToken queries the "verification_token" edge of the User entity.
func (u *User) QueryVerificationToken() *VerificationTokenQuery {
	return NewUserClient(u.config).QueryVerificationToken(u)
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return NewUserClient(u.config).QueryPosts(u)
}

// QueryLikes queries the "likes" edge of the User entity.
func (u *User) QueryLikes() *LikeQuery {
	return NewUserClient(u.config).QueryLikes(u)
}

// QuerySubscriptions queries the "subscriptions" edge of the User entity.
func (u *User) QuerySubscriptions() *SubscriptionQuery {
	return NewUserClient(u.config).QuerySubscriptions(u)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return NewUserClient(u.config).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return NewUserClient(u.config).QueryFollowing(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("profile_image_url=")
	builder.WriteString(u.ProfileImageURL)
	builder.WriteString(", ")
	builder.WriteString("stripe_customer_id=")
	builder.WriteString(u.StripeCustomerID)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(u.URL)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	if v := u.EmailVerified; v != nil {
		builder.WriteString("email_verified=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("bio=")
	builder.WriteString(u.Bio)
	builder.WriteString(", ")
	if v := u.Dob; v != nil {
		builder.WriteString("dob=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

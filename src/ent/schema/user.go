package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Default(uuid.NewString()),
		field.String("name"),
		field.String("username").Default(uuid.NewString()),
		field.String("url").Optional(),
		field.String("email").Unique(),
		field.Text("bio").Optional(),
		field.String("profile_image_url").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Has-many relationship to post
		edge.To("posts", Post.Type),

		// Has-many relationship to like
		edge.To("likes", Like.Type),

		// Has-many relationship to subscriptions
		edge.To("subscriptions", Subscription.Type),

		// Many-to-many relationship between followers and followings
		edge.To("following", User.Type).From("followers"),
	}
}
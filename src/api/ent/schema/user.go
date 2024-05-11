package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "user-" + cuid.New()
			}).
			NotEmpty().
			Unique().
			Immutable(),
		field.String("name"),
		field.String("username").Default(cuid.New()),
		field.String("profile_image_url").Optional(),
		field.String("stripe_customer_id").Optional(),
		field.String("password").Sensitive(),
		field.String("url").Optional(),
		field.String("email").Unique(),
		field.Bool("email_verified").Default(false).Nillable(),
		field.Text("bio").Optional(),
		field.String("dob").Optional().Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//
		edge.To("verification_token", VerificationToken.Type).Unique().Annotations(entsql.OnDelete(entsql.Cascade)),

		// Has-many relationship to post
		edge.To("posts", Post.Type),

		// Has-many relationship to like
		edge.To("likes", Like.Type).Annotations(entsql.OnDelete(entsql.Cascade)),

		// Has-many relationship to subscriptions
		edge.To("subscriptions", Subscription.Type).Annotations(entsql.OnDelete(entsql.Cascade)),

		// Many-to-many relationship between followers and followings
		edge.To("following", User.Type).From("followers"),
	}
}

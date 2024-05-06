package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "sub-" + cuid.New()
			}).
			NotEmpty().
			Unique().
			Immutable(),
		field.String("name"),
		field.Text("description").Nillable(),
		field.Int("price").Nillable(),
		field.Int("trial_period_days").Nillable(),
		field.Bool("is_archived").Default(false).Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		// Belongs-to relationship to a user
		edge.From("user", User.Type).Ref("subscriptions").Unique(),

		// Has-many relationship to post
		edge.To("posts", Post.Type),
	}
}

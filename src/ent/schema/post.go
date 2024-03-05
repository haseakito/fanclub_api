package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("user_id"),
		field.String("title"),
		field.Text("description").Optional(),
		field.Int("price").Optional(),
		field.Bool("is_featured").Default(false).Nillable(),
		field.Bool("status").Default(false).Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		// Has-many relationship to category
		edge.To("categories", Category.Type),

		// Has-many relationship to asset
		edge.To("assets", Asset.Type),

		// Belongs-to relationship to subscription
		edge.From("subscriptions", Subscription.Type).Ref("posts"),
	}
}

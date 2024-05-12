package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "order-" + cuid.New()
			}).
			NotEmpty().
			Unique().
			Immutable(),
		field.Int64("amount"),
		field.Enum("status").
			Values("pending", "completed", "processing", "canceled").
			Default("pending"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		// Belongs-to relationship to a post
		edge.From("post", Post.Type).
			Ref("orders").
			Unique(),

		// Belongs-to relationship to a user
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
	}
}

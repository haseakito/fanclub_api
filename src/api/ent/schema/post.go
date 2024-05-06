package schema

import (	
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "post-" + cuid.New()
			}).
			NotEmpty().
			Unique().
			Immutable(),
		field.String("title"),
		field.Text("description").Optional(),
		field.String("thumbnail_url").Optional(),
		field.String("video_url").Optional(),
		field.String("mux_asset_id").Optional(),
		field.String("mux_playback_id").Optional(),
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
		// Belongs-to relationship to a user
		edge.From("user", User.Type).Ref("posts").Unique(),

		// Belongs-to relationship to subscriptions
		edge.From("subscriptions", Subscription.Type).Ref("posts"),

		// Has-many relationship to like
		edge.To("likes", Like.Type),

		// Has-many relationship to category
		edge.To("categories", Category.Type),
	}
}

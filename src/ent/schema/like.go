package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("post_id", uuid.UUID{}).Default(uuid.New).Optional(),
		field.String("user_id").Default(uuid.NewString()).Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		// Belongs-to relationship to a post
		edge.From("post", Post.Type).
			Ref("likes").
			Field("post_id").
			Unique(),

		// Belongs-to relationship to a user
		edge.From("user", User.Type).
			Ref("likes").
			Field("user_id").
			Unique(),
	}
}

// Indexes of the Like.
func (Like) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id", "user_id"),
	}
}

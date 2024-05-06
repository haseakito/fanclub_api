package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"	
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		// Belongs-to relationship to a post
		edge.From("post", Post.Type).
			Ref("likes").			
			Unique(),

		// Belongs-to relationship to a user
		edge.From("user", User.Type).
			Ref("likes").			
			Unique(),
	}
}

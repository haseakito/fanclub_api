package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VerificationToken holds the schema definition for the VerificationToken entity.
type VerificationToken struct {
	ent.Schema
}

// Fields of the VerificationToken.
func (VerificationToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("verification_code").NotEmpty(),
		field.Time("expires_at"),
	}
}

// Edges of the VerificationToken.
func (VerificationToken) Edges() []ent.Edge {
	return []ent.Edge{
		// Belongs-to relationship to a user
		edge.From("user", User.Type).
			Ref("verification_token").
			Unique().
			Required(),
	}
}

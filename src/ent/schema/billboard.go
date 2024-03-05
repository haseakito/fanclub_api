package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Billboard holds the schema definition for the Billboard entity.
type Billboard struct {
	ent.Schema
}

// Fields of the Billboard.
func (Billboard) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title"),
		field.Text("description").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Billboard.
func (Billboard) Edges() []ent.Edge {
	return []ent.Edge{
		// Has-one relationship to asset
		edge.To("asset", Asset.Type).Unique(),
	}
}

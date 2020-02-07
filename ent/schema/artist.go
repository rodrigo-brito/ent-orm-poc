package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Artist struct {
	ent.Schema
}

func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Unique(),
		field.Int("age").Positive(),
		field.String("phone").Optional().Nillable(),
		field.Int64("plays").NonNegative(),
		field.Enum("gender").Values("f", "m"),
	}
}

func (Artist) Config() ent.Config {
	return ent.Config{
		Table: "Artistas",
	}
}

func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("songs", Song.Type),
	}
}

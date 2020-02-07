package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type Song struct {
	ent.Schema
}

func (Song) Fields() []ent.Field {
	return []ent.Field{
		field.Int("name").Positive(),
		field.String("phone").Optional().Nillable(),
		field.Int64("plays").NonNegative(),
		field.Enum("gender").Values("f", "m"),
	}
}

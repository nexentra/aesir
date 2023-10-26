package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Eval holds the schema definition for the Eval entity.
type Eval struct {
	ent.Schema
}

// Fields of the Eval.
func (Eval) Fields() []ent.Field {
	return []ent.Field{
		field.String("snippet"),
		field.JSON("result", []string{}),
		field.String("time"),
	}
}

// Edges of the Eval.
func (Eval) Edges() []ent.Edge {
	return nil
}

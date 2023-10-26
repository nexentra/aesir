package schema

import (
"entgo.io/ent"
"entgo.io/ent/schema/field"
// "entgo.io/ent/schema/edge"
// "entgo.io/contrib/entgql"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").
			NotEmpty(),
		field.Bool("done").
			Default(false),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EvalsColumns holds the columns for the "evals" table.
	EvalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "snippet", Type: field.TypeString},
		{Name: "result", Type: field.TypeJSON},
		{Name: "time", Type: field.TypeString},
	}
	// EvalsTable holds the schema information for the "evals" table.
	EvalsTable = &schema.Table{
		Name:       "evals",
		Columns:    EvalsColumns,
		PrimaryKey: []*schema.Column{EvalsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EvalsTable,
	}
)

func init() {
}
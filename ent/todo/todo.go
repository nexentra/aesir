// Code generated by ent, DO NOT EDIT.

package todo

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldDone holds the string denoting the done field in the database.
	FieldDone = "done"
	// Table holds the table name of the todo in the database.
	Table = "todos"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldDone,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultDone holds the default value on creation for the "done" field.
	DefaultDone bool
)

// OrderOption defines the ordering options for the Todo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByDone orders the results by the done field.
func ByDone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDone, opts...).ToFunc()
}

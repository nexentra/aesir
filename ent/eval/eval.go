// Code generated by ent, DO NOT EDIT.

package eval

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the eval type in the database.
	Label = "eval"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSnippet holds the string denoting the snippet field in the database.
	FieldSnippet = "snippet"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// Table holds the table name of the eval in the database.
	Table = "evals"
)

// Columns holds all SQL columns for eval fields.
var Columns = []string{
	FieldID,
	FieldSnippet,
	FieldResult,
	FieldTime,
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

// OrderOption defines the ordering options for the Eval queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySnippet orders the results by the snippet field.
func BySnippet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnippet, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}
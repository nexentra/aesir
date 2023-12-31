// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nexentra/aesir/ent/eval"
)

// Eval is the model entity for the Eval schema.
type Eval struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Snippet holds the value of the "snippet" field.
	Snippet string `json:"snippet,omitempty"`
	// Result holds the value of the "result" field.
	Result []string `json:"result,omitempty"`
	// Time holds the value of the "time" field.
	Time         string `json:"time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Eval) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case eval.FieldResult:
			values[i] = new([]byte)
		case eval.FieldID:
			values[i] = new(sql.NullInt64)
		case eval.FieldSnippet, eval.FieldTime:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Eval fields.
func (e *Eval) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eval.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case eval.FieldSnippet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field snippet", values[i])
			} else if value.Valid {
				e.Snippet = value.String
			}
		case eval.FieldResult:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Result); err != nil {
					return fmt.Errorf("unmarshal field result: %w", err)
				}
			}
		case eval.FieldTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				e.Time = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Eval.
// This includes values selected through modifiers, order, etc.
func (e *Eval) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Eval.
// Note that you need to call Eval.Unwrap() before calling this method if this Eval
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Eval) Update() *EvalUpdateOne {
	return NewEvalClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Eval entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Eval) Unwrap() *Eval {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Eval is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Eval) String() string {
	var builder strings.Builder
	builder.WriteString("Eval(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("snippet=")
	builder.WriteString(e.Snippet)
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", e.Result))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(e.Time)
	builder.WriteByte(')')
	return builder.String()
}

// Evals is a parsable slice of Eval.
type Evals []*Eval

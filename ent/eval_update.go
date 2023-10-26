// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/nexentra/aesir/ent/eval"
	"github.com/nexentra/aesir/ent/predicate"
)

// EvalUpdate is the builder for updating Eval entities.
type EvalUpdate struct {
	config
	hooks    []Hook
	mutation *EvalMutation
}

// Where appends a list predicates to the EvalUpdate builder.
func (eu *EvalUpdate) Where(ps ...predicate.Eval) *EvalUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetSnippet sets the "snippet" field.
func (eu *EvalUpdate) SetSnippet(s string) *EvalUpdate {
	eu.mutation.SetSnippet(s)
	return eu
}

// SetResult sets the "result" field.
func (eu *EvalUpdate) SetResult(s []string) *EvalUpdate {
	eu.mutation.SetResult(s)
	return eu
}

// AppendResult appends s to the "result" field.
func (eu *EvalUpdate) AppendResult(s []string) *EvalUpdate {
	eu.mutation.AppendResult(s)
	return eu
}

// SetTime sets the "time" field.
func (eu *EvalUpdate) SetTime(s string) *EvalUpdate {
	eu.mutation.SetTime(s)
	return eu
}

// Mutation returns the EvalMutation object of the builder.
func (eu *EvalUpdate) Mutation() *EvalMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EvalUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EvalUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EvalUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EvalUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EvalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(eval.Table, eval.Columns, sqlgraph.NewFieldSpec(eval.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Snippet(); ok {
		_spec.SetField(eval.FieldSnippet, field.TypeString, value)
	}
	if value, ok := eu.mutation.Result(); ok {
		_spec.SetField(eval.FieldResult, field.TypeJSON, value)
	}
	if value, ok := eu.mutation.AppendedResult(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, eval.FieldResult, value)
		})
	}
	if value, ok := eu.mutation.Time(); ok {
		_spec.SetField(eval.FieldTime, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eval.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EvalUpdateOne is the builder for updating a single Eval entity.
type EvalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EvalMutation
}

// SetSnippet sets the "snippet" field.
func (euo *EvalUpdateOne) SetSnippet(s string) *EvalUpdateOne {
	euo.mutation.SetSnippet(s)
	return euo
}

// SetResult sets the "result" field.
func (euo *EvalUpdateOne) SetResult(s []string) *EvalUpdateOne {
	euo.mutation.SetResult(s)
	return euo
}

// AppendResult appends s to the "result" field.
func (euo *EvalUpdateOne) AppendResult(s []string) *EvalUpdateOne {
	euo.mutation.AppendResult(s)
	return euo
}

// SetTime sets the "time" field.
func (euo *EvalUpdateOne) SetTime(s string) *EvalUpdateOne {
	euo.mutation.SetTime(s)
	return euo
}

// Mutation returns the EvalMutation object of the builder.
func (euo *EvalUpdateOne) Mutation() *EvalMutation {
	return euo.mutation
}

// Where appends a list predicates to the EvalUpdate builder.
func (euo *EvalUpdateOne) Where(ps ...predicate.Eval) *EvalUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EvalUpdateOne) Select(field string, fields ...string) *EvalUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Eval entity.
func (euo *EvalUpdateOne) Save(ctx context.Context) (*Eval, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EvalUpdateOne) SaveX(ctx context.Context) *Eval {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EvalUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EvalUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EvalUpdateOne) sqlSave(ctx context.Context) (_node *Eval, err error) {
	_spec := sqlgraph.NewUpdateSpec(eval.Table, eval.Columns, sqlgraph.NewFieldSpec(eval.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Eval.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eval.FieldID)
		for _, f := range fields {
			if !eval.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eval.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Snippet(); ok {
		_spec.SetField(eval.FieldSnippet, field.TypeString, value)
	}
	if value, ok := euo.mutation.Result(); ok {
		_spec.SetField(eval.FieldResult, field.TypeJSON, value)
	}
	if value, ok := euo.mutation.AppendedResult(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, eval.FieldResult, value)
		})
	}
	if value, ok := euo.mutation.Time(); ok {
		_spec.SetField(eval.FieldTime, field.TypeString, value)
	}
	_node = &Eval{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eval.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
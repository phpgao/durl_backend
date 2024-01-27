// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/phpgao/durl_backend/internal/data/ent/predicate"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
)

// TLeafDelete is the builder for deleting a TLeaf entity.
type TLeafDelete struct {
	config
	hooks    []Hook
	mutation *TLeafMutation
}

// Where appends a list predicates to the TLeafDelete builder.
func (td *TLeafDelete) Where(ps ...predicate.TLeaf) *TLeafDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TLeafDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TLeafDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TLeafDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tleaf.Table, sqlgraph.NewFieldSpec(tleaf.FieldID, field.TypeInt64))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TLeafDeleteOne is the builder for deleting a single TLeaf entity.
type TLeafDeleteOne struct {
	td *TLeafDelete
}

// Where appends a list predicates to the TLeafDelete builder.
func (tdo *TLeafDeleteOne) Where(ps ...predicate.TLeaf) *TLeafDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TLeafDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tleaf.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TLeafDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}

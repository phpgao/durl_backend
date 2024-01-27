// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/phpgao/durl_backend/internal/data/ent/predicate"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
)

// TShortUrlDelete is the builder for deleting a TShortUrl entity.
type TShortUrlDelete struct {
	config
	hooks    []Hook
	mutation *TShortUrlMutation
}

// Where appends a list predicates to the TShortUrlDelete builder.
func (tud *TShortUrlDelete) Where(ps ...predicate.TShortUrl) *TShortUrlDelete {
	tud.mutation.Where(ps...)
	return tud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tud *TShortUrlDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tud.sqlExec, tud.mutation, tud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tud *TShortUrlDelete) ExecX(ctx context.Context) int {
	n, err := tud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tud *TShortUrlDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tshorturl.Table, sqlgraph.NewFieldSpec(tshorturl.FieldID, field.TypeInt64))
	if ps := tud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tud.mutation.done = true
	return affected, err
}

// TShortUrlDeleteOne is the builder for deleting a single TShortUrl entity.
type TShortUrlDeleteOne struct {
	tud *TShortUrlDelete
}

// Where appends a list predicates to the TShortUrlDelete builder.
func (tudo *TShortUrlDeleteOne) Where(ps ...predicate.TShortUrl) *TShortUrlDeleteOne {
	tudo.tud.mutation.Where(ps...)
	return tudo
}

// Exec executes the deletion query.
func (tudo *TShortUrlDeleteOne) Exec(ctx context.Context) error {
	n, err := tudo.tud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tshorturl.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tudo *TShortUrlDeleteOne) ExecX(ctx context.Context) {
	if err := tudo.Exec(ctx); err != nil {
		panic(err)
	}
}

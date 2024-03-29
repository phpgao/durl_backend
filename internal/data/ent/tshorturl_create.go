// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
)

// TShortUrlCreate is the builder for creating a TShortUrl entity.
type TShortUrlCreate struct {
	config
	mutation *TShortUrlMutation
	hooks    []Hook
}

// SetBizID sets the "biz_id" field.
func (tuc *TShortUrlCreate) SetBizID(i int64) *TShortUrlCreate {
	tuc.mutation.SetBizID(i)
	return tuc
}

// SetOrigin sets the "origin" field.
func (tuc *TShortUrlCreate) SetOrigin(s string) *TShortUrlCreate {
	tuc.mutation.SetOrigin(s)
	return tuc
}

// SetShort sets the "short" field.
func (tuc *TShortUrlCreate) SetShort(i int64) *TShortUrlCreate {
	tuc.mutation.SetShort(i)
	return tuc
}

// SetVisit sets the "visit" field.
func (tuc *TShortUrlCreate) SetVisit(i int64) *TShortUrlCreate {
	tuc.mutation.SetVisit(i)
	return tuc
}

// SetNillableVisit sets the "visit" field if the given value is not nil.
func (tuc *TShortUrlCreate) SetNillableVisit(i *int64) *TShortUrlCreate {
	if i != nil {
		tuc.SetVisit(*i)
	}
	return tuc
}

// SetCreatedAt sets the "created_at" field.
func (tuc *TShortUrlCreate) SetCreatedAt(i int64) *TShortUrlCreate {
	tuc.mutation.SetCreatedAt(i)
	return tuc
}

// SetUpdatedAt sets the "updated_at" field.
func (tuc *TShortUrlCreate) SetUpdatedAt(i int64) *TShortUrlCreate {
	tuc.mutation.SetUpdatedAt(i)
	return tuc
}

// SetExpiredAt sets the "expired_at" field.
func (tuc *TShortUrlCreate) SetExpiredAt(i int64) *TShortUrlCreate {
	tuc.mutation.SetExpiredAt(i)
	return tuc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tuc *TShortUrlCreate) SetNillableExpiredAt(i *int64) *TShortUrlCreate {
	if i != nil {
		tuc.SetExpiredAt(*i)
	}
	return tuc
}

// SetID sets the "id" field.
func (tuc *TShortUrlCreate) SetID(i int64) *TShortUrlCreate {
	tuc.mutation.SetID(i)
	return tuc
}

// Mutation returns the TShortUrlMutation object of the builder.
func (tuc *TShortUrlCreate) Mutation() *TShortUrlMutation {
	return tuc.mutation
}

// Save creates the TShortUrl in the database.
func (tuc *TShortUrlCreate) Save(ctx context.Context) (*TShortUrl, error) {
	tuc.defaults()
	return withHooks(ctx, tuc.sqlSave, tuc.mutation, tuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tuc *TShortUrlCreate) SaveX(ctx context.Context) *TShortUrl {
	v, err := tuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tuc *TShortUrlCreate) Exec(ctx context.Context) error {
	_, err := tuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuc *TShortUrlCreate) ExecX(ctx context.Context) {
	if err := tuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuc *TShortUrlCreate) defaults() {
	if _, ok := tuc.mutation.Visit(); !ok {
		v := tshorturl.DefaultVisit
		tuc.mutation.SetVisit(v)
	}
	if _, ok := tuc.mutation.ExpiredAt(); !ok {
		v := tshorturl.DefaultExpiredAt
		tuc.mutation.SetExpiredAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuc *TShortUrlCreate) check() error {
	if _, ok := tuc.mutation.BizID(); !ok {
		return &ValidationError{Name: "biz_id", err: errors.New(`ent: missing required field "TShortUrl.biz_id"`)}
	}
	if _, ok := tuc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "TShortUrl.origin"`)}
	}
	if v, ok := tuc.mutation.Origin(); ok {
		if err := tshorturl.OriginValidator(v); err != nil {
			return &ValidationError{Name: "origin", err: fmt.Errorf(`ent: validator failed for field "TShortUrl.origin": %w`, err)}
		}
	}
	if _, ok := tuc.mutation.Short(); !ok {
		return &ValidationError{Name: "short", err: errors.New(`ent: missing required field "TShortUrl.short"`)}
	}
	if _, ok := tuc.mutation.Visit(); !ok {
		return &ValidationError{Name: "visit", err: errors.New(`ent: missing required field "TShortUrl.visit"`)}
	}
	if _, ok := tuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TShortUrl.created_at"`)}
	}
	if _, ok := tuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TShortUrl.updated_at"`)}
	}
	if _, ok := tuc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New(`ent: missing required field "TShortUrl.expired_at"`)}
	}
	return nil
}

func (tuc *TShortUrlCreate) sqlSave(ctx context.Context) (*TShortUrl, error) {
	if err := tuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tuc.mutation.id = &_node.ID
	tuc.mutation.done = true
	return _node, nil
}

func (tuc *TShortUrlCreate) createSpec() (*TShortUrl, *sqlgraph.CreateSpec) {
	var (
		_node = &TShortUrl{config: tuc.config}
		_spec = sqlgraph.NewCreateSpec(tshorturl.Table, sqlgraph.NewFieldSpec(tshorturl.FieldID, field.TypeInt64))
	)
	if id, ok := tuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tuc.mutation.BizID(); ok {
		_spec.SetField(tshorturl.FieldBizID, field.TypeInt64, value)
		_node.BizID = value
	}
	if value, ok := tuc.mutation.Origin(); ok {
		_spec.SetField(tshorturl.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := tuc.mutation.Short(); ok {
		_spec.SetField(tshorturl.FieldShort, field.TypeInt64, value)
		_node.Short = value
	}
	if value, ok := tuc.mutation.Visit(); ok {
		_spec.SetField(tshorturl.FieldVisit, field.TypeInt64, value)
		_node.Visit = value
	}
	if value, ok := tuc.mutation.CreatedAt(); ok {
		_spec.SetField(tshorturl.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := tuc.mutation.UpdatedAt(); ok {
		_spec.SetField(tshorturl.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := tuc.mutation.ExpiredAt(); ok {
		_spec.SetField(tshorturl.FieldExpiredAt, field.TypeInt64, value)
		_node.ExpiredAt = value
	}
	return _node, _spec
}

// TShortUrlCreateBulk is the builder for creating many TShortUrl entities in bulk.
type TShortUrlCreateBulk struct {
	config
	err      error
	builders []*TShortUrlCreate
}

// Save creates the TShortUrl entities in the database.
func (tucb *TShortUrlCreateBulk) Save(ctx context.Context) ([]*TShortUrl, error) {
	if tucb.err != nil {
		return nil, tucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tucb.builders))
	nodes := make([]*TShortUrl, len(tucb.builders))
	mutators := make([]Mutator, len(tucb.builders))
	for i := range tucb.builders {
		func(i int, root context.Context) {
			builder := tucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TShortUrlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tucb *TShortUrlCreateBulk) SaveX(ctx context.Context) []*TShortUrl {
	v, err := tucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tucb *TShortUrlCreateBulk) Exec(ctx context.Context) error {
	_, err := tucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tucb *TShortUrlCreateBulk) ExecX(ctx context.Context) {
	if err := tucb.Exec(ctx); err != nil {
		panic(err)
	}
}

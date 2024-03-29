// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
)

// TLeafCreate is the builder for creating a TLeaf entity.
type TLeafCreate struct {
	config
	mutation *TLeafMutation
	hooks    []Hook
}

// SetBizTag sets the "biz_tag" field.
func (tc *TLeafCreate) SetBizTag(s string) *TLeafCreate {
	tc.mutation.SetBizTag(s)
	return tc
}

// SetMaxID sets the "max_id" field.
func (tc *TLeafCreate) SetMaxID(i int64) *TLeafCreate {
	tc.mutation.SetMaxID(i)
	return tc
}

// SetNillableMaxID sets the "max_id" field if the given value is not nil.
func (tc *TLeafCreate) SetNillableMaxID(i *int64) *TLeafCreate {
	if i != nil {
		tc.SetMaxID(*i)
	}
	return tc
}

// SetStep sets the "step" field.
func (tc *TLeafCreate) SetStep(i int64) *TLeafCreate {
	tc.mutation.SetStep(i)
	return tc
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tc *TLeafCreate) SetNillableStep(i *int64) *TLeafCreate {
	if i != nil {
		tc.SetStep(*i)
	}
	return tc
}

// SetDesc sets the "desc" field.
func (tc *TLeafCreate) SetDesc(s string) *TLeafCreate {
	tc.mutation.SetDesc(s)
	return tc
}

// SetVersion sets the "version" field.
func (tc *TLeafCreate) SetVersion(i int32) *TLeafCreate {
	tc.mutation.SetVersion(i)
	return tc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tc *TLeafCreate) SetNillableVersion(i *int32) *TLeafCreate {
	if i != nil {
		tc.SetVersion(*i)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TLeafCreate) SetCreatedAt(i int64) *TLeafCreate {
	tc.mutation.SetCreatedAt(i)
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TLeafCreate) SetUpdatedAt(i int64) *TLeafCreate {
	tc.mutation.SetUpdatedAt(i)
	return tc
}

// SetID sets the "id" field.
func (tc *TLeafCreate) SetID(i int64) *TLeafCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TLeafMutation object of the builder.
func (tc *TLeafCreate) Mutation() *TLeafMutation {
	return tc.mutation
}

// Save creates the TLeaf in the database.
func (tc *TLeafCreate) Save(ctx context.Context) (*TLeaf, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TLeafCreate) SaveX(ctx context.Context) *TLeaf {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TLeafCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TLeafCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TLeafCreate) defaults() {
	if _, ok := tc.mutation.MaxID(); !ok {
		v := tleaf.DefaultMaxID
		tc.mutation.SetMaxID(v)
	}
	if _, ok := tc.mutation.Step(); !ok {
		v := tleaf.DefaultStep
		tc.mutation.SetStep(v)
	}
	if _, ok := tc.mutation.Version(); !ok {
		v := tleaf.DefaultVersion
		tc.mutation.SetVersion(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TLeafCreate) check() error {
	if _, ok := tc.mutation.BizTag(); !ok {
		return &ValidationError{Name: "biz_tag", err: errors.New(`ent: missing required field "TLeaf.biz_tag"`)}
	}
	if v, ok := tc.mutation.BizTag(); ok {
		if err := tleaf.BizTagValidator(v); err != nil {
			return &ValidationError{Name: "biz_tag", err: fmt.Errorf(`ent: validator failed for field "TLeaf.biz_tag": %w`, err)}
		}
	}
	if _, ok := tc.mutation.MaxID(); !ok {
		return &ValidationError{Name: "max_id", err: errors.New(`ent: missing required field "TLeaf.max_id"`)}
	}
	if _, ok := tc.mutation.Step(); !ok {
		return &ValidationError{Name: "step", err: errors.New(`ent: missing required field "TLeaf.step"`)}
	}
	if _, ok := tc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "TLeaf.desc"`)}
	}
	if v, ok := tc.mutation.Desc(); ok {
		if err := tleaf.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf(`ent: validator failed for field "TLeaf.desc": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "TLeaf.version"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TLeaf.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TLeaf.updated_at"`)}
	}
	return nil
}

func (tc *TLeafCreate) sqlSave(ctx context.Context) (*TLeaf, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TLeafCreate) createSpec() (*TLeaf, *sqlgraph.CreateSpec) {
	var (
		_node = &TLeaf{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tleaf.Table, sqlgraph.NewFieldSpec(tleaf.FieldID, field.TypeInt64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.BizTag(); ok {
		_spec.SetField(tleaf.FieldBizTag, field.TypeString, value)
		_node.BizTag = value
	}
	if value, ok := tc.mutation.MaxID(); ok {
		_spec.SetField(tleaf.FieldMaxID, field.TypeInt64, value)
		_node.MaxID = value
	}
	if value, ok := tc.mutation.Step(); ok {
		_spec.SetField(tleaf.FieldStep, field.TypeInt64, value)
		_node.Step = value
	}
	if value, ok := tc.mutation.Desc(); ok {
		_spec.SetField(tleaf.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := tc.mutation.Version(); ok {
		_spec.SetField(tleaf.FieldVersion, field.TypeInt32, value)
		_node.Version = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tleaf.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tleaf.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// TLeafCreateBulk is the builder for creating many TLeaf entities in bulk.
type TLeafCreateBulk struct {
	config
	err      error
	builders []*TLeafCreate
}

// Save creates the TLeaf entities in the database.
func (tcb *TLeafCreateBulk) Save(ctx context.Context) ([]*TLeaf, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*TLeaf, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TLeafMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TLeafCreateBulk) SaveX(ctx context.Context) []*TLeaf {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TLeafCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TLeafCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

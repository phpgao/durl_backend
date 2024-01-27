// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/phpgao/durl_backend/internal/data/ent/predicate"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
)

// TLeafUpdate is the builder for updating TLeaf entities.
type TLeafUpdate struct {
	config
	hooks     []Hook
	mutation  *TLeafMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TLeafUpdate builder.
func (tu *TLeafUpdate) Where(ps ...predicate.TLeaf) *TLeafUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetBizTag sets the "biz_tag" field.
func (tu *TLeafUpdate) SetBizTag(s string) *TLeafUpdate {
	tu.mutation.SetBizTag(s)
	return tu
}

// SetNillableBizTag sets the "biz_tag" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableBizTag(s *string) *TLeafUpdate {
	if s != nil {
		tu.SetBizTag(*s)
	}
	return tu
}

// SetMaxID sets the "max_id" field.
func (tu *TLeafUpdate) SetMaxID(i int64) *TLeafUpdate {
	tu.mutation.ResetMaxID()
	tu.mutation.SetMaxID(i)
	return tu
}

// SetNillableMaxID sets the "max_id" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableMaxID(i *int64) *TLeafUpdate {
	if i != nil {
		tu.SetMaxID(*i)
	}
	return tu
}

// AddMaxID adds i to the "max_id" field.
func (tu *TLeafUpdate) AddMaxID(i int64) *TLeafUpdate {
	tu.mutation.AddMaxID(i)
	return tu
}

// SetStep sets the "step" field.
func (tu *TLeafUpdate) SetStep(i int64) *TLeafUpdate {
	tu.mutation.ResetStep()
	tu.mutation.SetStep(i)
	return tu
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableStep(i *int64) *TLeafUpdate {
	if i != nil {
		tu.SetStep(*i)
	}
	return tu
}

// AddStep adds i to the "step" field.
func (tu *TLeafUpdate) AddStep(i int64) *TLeafUpdate {
	tu.mutation.AddStep(i)
	return tu
}

// SetDesc sets the "desc" field.
func (tu *TLeafUpdate) SetDesc(s string) *TLeafUpdate {
	tu.mutation.SetDesc(s)
	return tu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableDesc(s *string) *TLeafUpdate {
	if s != nil {
		tu.SetDesc(*s)
	}
	return tu
}

// SetVersion sets the "version" field.
func (tu *TLeafUpdate) SetVersion(i int32) *TLeafUpdate {
	tu.mutation.ResetVersion()
	tu.mutation.SetVersion(i)
	return tu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableVersion(i *int32) *TLeafUpdate {
	if i != nil {
		tu.SetVersion(*i)
	}
	return tu
}

// AddVersion adds i to the "version" field.
func (tu *TLeafUpdate) AddVersion(i int32) *TLeafUpdate {
	tu.mutation.AddVersion(i)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TLeafUpdate) SetUpdatedAt(i int64) *TLeafUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(i)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TLeafUpdate) SetNillableUpdatedAt(i *int64) *TLeafUpdate {
	if i != nil {
		tu.SetUpdatedAt(*i)
	}
	return tu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tu *TLeafUpdate) AddUpdatedAt(i int64) *TLeafUpdate {
	tu.mutation.AddUpdatedAt(i)
	return tu
}

// Mutation returns the TLeafMutation object of the builder.
func (tu *TLeafUpdate) Mutation() *TLeafMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TLeafUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TLeafUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TLeafUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TLeafUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TLeafUpdate) check() error {
	if v, ok := tu.mutation.BizTag(); ok {
		if err := tleaf.BizTagValidator(v); err != nil {
			return &ValidationError{Name: "biz_tag", err: fmt.Errorf(`ent: validator failed for field "TLeaf.biz_tag": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Desc(); ok {
		if err := tleaf.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf(`ent: validator failed for field "TLeaf.desc": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TLeafUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TLeafUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TLeafUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tleaf.Table, tleaf.Columns, sqlgraph.NewFieldSpec(tleaf.FieldID, field.TypeInt64))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.BizTag(); ok {
		_spec.SetField(tleaf.FieldBizTag, field.TypeString, value)
	}
	if value, ok := tu.mutation.MaxID(); ok {
		_spec.SetField(tleaf.FieldMaxID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedMaxID(); ok {
		_spec.AddField(tleaf.FieldMaxID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.Step(); ok {
		_spec.SetField(tleaf.FieldStep, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedStep(); ok {
		_spec.AddField(tleaf.FieldStep, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.Desc(); ok {
		_spec.SetField(tleaf.FieldDesc, field.TypeString, value)
	}
	if value, ok := tu.mutation.Version(); ok {
		_spec.SetField(tleaf.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.AddedVersion(); ok {
		_spec.AddField(tleaf.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tleaf.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tleaf.FieldUpdatedAt, field.TypeInt64, value)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tleaf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TLeafUpdateOne is the builder for updating a single TLeaf entity.
type TLeafUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TLeafMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetBizTag sets the "biz_tag" field.
func (tuo *TLeafUpdateOne) SetBizTag(s string) *TLeafUpdateOne {
	tuo.mutation.SetBizTag(s)
	return tuo
}

// SetNillableBizTag sets the "biz_tag" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableBizTag(s *string) *TLeafUpdateOne {
	if s != nil {
		tuo.SetBizTag(*s)
	}
	return tuo
}

// SetMaxID sets the "max_id" field.
func (tuo *TLeafUpdateOne) SetMaxID(i int64) *TLeafUpdateOne {
	tuo.mutation.ResetMaxID()
	tuo.mutation.SetMaxID(i)
	return tuo
}

// SetNillableMaxID sets the "max_id" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableMaxID(i *int64) *TLeafUpdateOne {
	if i != nil {
		tuo.SetMaxID(*i)
	}
	return tuo
}

// AddMaxID adds i to the "max_id" field.
func (tuo *TLeafUpdateOne) AddMaxID(i int64) *TLeafUpdateOne {
	tuo.mutation.AddMaxID(i)
	return tuo
}

// SetStep sets the "step" field.
func (tuo *TLeafUpdateOne) SetStep(i int64) *TLeafUpdateOne {
	tuo.mutation.ResetStep()
	tuo.mutation.SetStep(i)
	return tuo
}

// SetNillableStep sets the "step" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableStep(i *int64) *TLeafUpdateOne {
	if i != nil {
		tuo.SetStep(*i)
	}
	return tuo
}

// AddStep adds i to the "step" field.
func (tuo *TLeafUpdateOne) AddStep(i int64) *TLeafUpdateOne {
	tuo.mutation.AddStep(i)
	return tuo
}

// SetDesc sets the "desc" field.
func (tuo *TLeafUpdateOne) SetDesc(s string) *TLeafUpdateOne {
	tuo.mutation.SetDesc(s)
	return tuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableDesc(s *string) *TLeafUpdateOne {
	if s != nil {
		tuo.SetDesc(*s)
	}
	return tuo
}

// SetVersion sets the "version" field.
func (tuo *TLeafUpdateOne) SetVersion(i int32) *TLeafUpdateOne {
	tuo.mutation.ResetVersion()
	tuo.mutation.SetVersion(i)
	return tuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableVersion(i *int32) *TLeafUpdateOne {
	if i != nil {
		tuo.SetVersion(*i)
	}
	return tuo
}

// AddVersion adds i to the "version" field.
func (tuo *TLeafUpdateOne) AddVersion(i int32) *TLeafUpdateOne {
	tuo.mutation.AddVersion(i)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TLeafUpdateOne) SetUpdatedAt(i int64) *TLeafUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(i)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TLeafUpdateOne) SetNillableUpdatedAt(i *int64) *TLeafUpdateOne {
	if i != nil {
		tuo.SetUpdatedAt(*i)
	}
	return tuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuo *TLeafUpdateOne) AddUpdatedAt(i int64) *TLeafUpdateOne {
	tuo.mutation.AddUpdatedAt(i)
	return tuo
}

// Mutation returns the TLeafMutation object of the builder.
func (tuo *TLeafUpdateOne) Mutation() *TLeafMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TLeafUpdate builder.
func (tuo *TLeafUpdateOne) Where(ps ...predicate.TLeaf) *TLeafUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TLeafUpdateOne) Select(field string, fields ...string) *TLeafUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated TLeaf entity.
func (tuo *TLeafUpdateOne) Save(ctx context.Context) (*TLeaf, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TLeafUpdateOne) SaveX(ctx context.Context) *TLeaf {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TLeafUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TLeafUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TLeafUpdateOne) check() error {
	if v, ok := tuo.mutation.BizTag(); ok {
		if err := tleaf.BizTagValidator(v); err != nil {
			return &ValidationError{Name: "biz_tag", err: fmt.Errorf(`ent: validator failed for field "TLeaf.biz_tag": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Desc(); ok {
		if err := tleaf.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf(`ent: validator failed for field "TLeaf.desc": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TLeafUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TLeafUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TLeafUpdateOne) sqlSave(ctx context.Context) (_node *TLeaf, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tleaf.Table, tleaf.Columns, sqlgraph.NewFieldSpec(tleaf.FieldID, field.TypeInt64))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TLeaf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tleaf.FieldID)
		for _, f := range fields {
			if !tleaf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tleaf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.BizTag(); ok {
		_spec.SetField(tleaf.FieldBizTag, field.TypeString, value)
	}
	if value, ok := tuo.mutation.MaxID(); ok {
		_spec.SetField(tleaf.FieldMaxID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedMaxID(); ok {
		_spec.AddField(tleaf.FieldMaxID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.Step(); ok {
		_spec.SetField(tleaf.FieldStep, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedStep(); ok {
		_spec.AddField(tleaf.FieldStep, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.Desc(); ok {
		_spec.SetField(tleaf.FieldDesc, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Version(); ok {
		_spec.SetField(tleaf.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.AddedVersion(); ok {
		_spec.AddField(tleaf.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tleaf.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tleaf.FieldUpdatedAt, field.TypeInt64, value)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &TLeaf{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tleaf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
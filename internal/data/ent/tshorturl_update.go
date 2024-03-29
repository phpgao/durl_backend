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
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
)

// TShortUrlUpdate is the builder for updating TShortUrl entities.
type TShortUrlUpdate struct {
	config
	hooks     []Hook
	mutation  *TShortUrlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TShortUrlUpdate builder.
func (tuu *TShortUrlUpdate) Where(ps ...predicate.TShortUrl) *TShortUrlUpdate {
	tuu.mutation.Where(ps...)
	return tuu
}

// SetBizID sets the "biz_id" field.
func (tuu *TShortUrlUpdate) SetBizID(i int64) *TShortUrlUpdate {
	tuu.mutation.ResetBizID()
	tuu.mutation.SetBizID(i)
	return tuu
}

// SetNillableBizID sets the "biz_id" field if the given value is not nil.
func (tuu *TShortUrlUpdate) SetNillableBizID(i *int64) *TShortUrlUpdate {
	if i != nil {
		tuu.SetBizID(*i)
	}
	return tuu
}

// AddBizID adds i to the "biz_id" field.
func (tuu *TShortUrlUpdate) AddBizID(i int64) *TShortUrlUpdate {
	tuu.mutation.AddBizID(i)
	return tuu
}

// SetOrigin sets the "origin" field.
func (tuu *TShortUrlUpdate) SetOrigin(s string) *TShortUrlUpdate {
	tuu.mutation.SetOrigin(s)
	return tuu
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (tuu *TShortUrlUpdate) SetNillableOrigin(s *string) *TShortUrlUpdate {
	if s != nil {
		tuu.SetOrigin(*s)
	}
	return tuu
}

// SetVisit sets the "visit" field.
func (tuu *TShortUrlUpdate) SetVisit(i int64) *TShortUrlUpdate {
	tuu.mutation.ResetVisit()
	tuu.mutation.SetVisit(i)
	return tuu
}

// SetNillableVisit sets the "visit" field if the given value is not nil.
func (tuu *TShortUrlUpdate) SetNillableVisit(i *int64) *TShortUrlUpdate {
	if i != nil {
		tuu.SetVisit(*i)
	}
	return tuu
}

// AddVisit adds i to the "visit" field.
func (tuu *TShortUrlUpdate) AddVisit(i int64) *TShortUrlUpdate {
	tuu.mutation.AddVisit(i)
	return tuu
}

// SetUpdatedAt sets the "updated_at" field.
func (tuu *TShortUrlUpdate) SetUpdatedAt(i int64) *TShortUrlUpdate {
	tuu.mutation.ResetUpdatedAt()
	tuu.mutation.SetUpdatedAt(i)
	return tuu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuu *TShortUrlUpdate) SetNillableUpdatedAt(i *int64) *TShortUrlUpdate {
	if i != nil {
		tuu.SetUpdatedAt(*i)
	}
	return tuu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuu *TShortUrlUpdate) AddUpdatedAt(i int64) *TShortUrlUpdate {
	tuu.mutation.AddUpdatedAt(i)
	return tuu
}

// SetExpiredAt sets the "expired_at" field.
func (tuu *TShortUrlUpdate) SetExpiredAt(i int64) *TShortUrlUpdate {
	tuu.mutation.ResetExpiredAt()
	tuu.mutation.SetExpiredAt(i)
	return tuu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tuu *TShortUrlUpdate) SetNillableExpiredAt(i *int64) *TShortUrlUpdate {
	if i != nil {
		tuu.SetExpiredAt(*i)
	}
	return tuu
}

// AddExpiredAt adds i to the "expired_at" field.
func (tuu *TShortUrlUpdate) AddExpiredAt(i int64) *TShortUrlUpdate {
	tuu.mutation.AddExpiredAt(i)
	return tuu
}

// Mutation returns the TShortUrlMutation object of the builder.
func (tuu *TShortUrlUpdate) Mutation() *TShortUrlMutation {
	return tuu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tuu *TShortUrlUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tuu.sqlSave, tuu.mutation, tuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuu *TShortUrlUpdate) SaveX(ctx context.Context) int {
	affected, err := tuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tuu *TShortUrlUpdate) Exec(ctx context.Context) error {
	_, err := tuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuu *TShortUrlUpdate) ExecX(ctx context.Context) {
	if err := tuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuu *TShortUrlUpdate) check() error {
	if v, ok := tuu.mutation.Origin(); ok {
		if err := tshorturl.OriginValidator(v); err != nil {
			return &ValidationError{Name: "origin", err: fmt.Errorf(`ent: validator failed for field "TShortUrl.origin": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuu *TShortUrlUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TShortUrlUpdate {
	tuu.modifiers = append(tuu.modifiers, modifiers...)
	return tuu
}

func (tuu *TShortUrlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tuu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tshorturl.Table, tshorturl.Columns, sqlgraph.NewFieldSpec(tshorturl.FieldID, field.TypeInt64))
	if ps := tuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuu.mutation.BizID(); ok {
		_spec.SetField(tshorturl.FieldBizID, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.AddedBizID(); ok {
		_spec.AddField(tshorturl.FieldBizID, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.Origin(); ok {
		_spec.SetField(tshorturl.FieldOrigin, field.TypeString, value)
	}
	if value, ok := tuu.mutation.Visit(); ok {
		_spec.SetField(tshorturl.FieldVisit, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.AddedVisit(); ok {
		_spec.AddField(tshorturl.FieldVisit, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.UpdatedAt(); ok {
		_spec.SetField(tshorturl.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tshorturl.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.ExpiredAt(); ok {
		_spec.SetField(tshorturl.FieldExpiredAt, field.TypeInt64, value)
	}
	if value, ok := tuu.mutation.AddedExpiredAt(); ok {
		_spec.AddField(tshorturl.FieldExpiredAt, field.TypeInt64, value)
	}
	_spec.AddModifiers(tuu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tshorturl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tuu.mutation.done = true
	return n, nil
}

// TShortUrlUpdateOne is the builder for updating a single TShortUrl entity.
type TShortUrlUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TShortUrlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetBizID sets the "biz_id" field.
func (tuuo *TShortUrlUpdateOne) SetBizID(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.ResetBizID()
	tuuo.mutation.SetBizID(i)
	return tuuo
}

// SetNillableBizID sets the "biz_id" field if the given value is not nil.
func (tuuo *TShortUrlUpdateOne) SetNillableBizID(i *int64) *TShortUrlUpdateOne {
	if i != nil {
		tuuo.SetBizID(*i)
	}
	return tuuo
}

// AddBizID adds i to the "biz_id" field.
func (tuuo *TShortUrlUpdateOne) AddBizID(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.AddBizID(i)
	return tuuo
}

// SetOrigin sets the "origin" field.
func (tuuo *TShortUrlUpdateOne) SetOrigin(s string) *TShortUrlUpdateOne {
	tuuo.mutation.SetOrigin(s)
	return tuuo
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (tuuo *TShortUrlUpdateOne) SetNillableOrigin(s *string) *TShortUrlUpdateOne {
	if s != nil {
		tuuo.SetOrigin(*s)
	}
	return tuuo
}

// SetVisit sets the "visit" field.
func (tuuo *TShortUrlUpdateOne) SetVisit(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.ResetVisit()
	tuuo.mutation.SetVisit(i)
	return tuuo
}

// SetNillableVisit sets the "visit" field if the given value is not nil.
func (tuuo *TShortUrlUpdateOne) SetNillableVisit(i *int64) *TShortUrlUpdateOne {
	if i != nil {
		tuuo.SetVisit(*i)
	}
	return tuuo
}

// AddVisit adds i to the "visit" field.
func (tuuo *TShortUrlUpdateOne) AddVisit(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.AddVisit(i)
	return tuuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuuo *TShortUrlUpdateOne) SetUpdatedAt(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.ResetUpdatedAt()
	tuuo.mutation.SetUpdatedAt(i)
	return tuuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuuo *TShortUrlUpdateOne) SetNillableUpdatedAt(i *int64) *TShortUrlUpdateOne {
	if i != nil {
		tuuo.SetUpdatedAt(*i)
	}
	return tuuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuuo *TShortUrlUpdateOne) AddUpdatedAt(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.AddUpdatedAt(i)
	return tuuo
}

// SetExpiredAt sets the "expired_at" field.
func (tuuo *TShortUrlUpdateOne) SetExpiredAt(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.ResetExpiredAt()
	tuuo.mutation.SetExpiredAt(i)
	return tuuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (tuuo *TShortUrlUpdateOne) SetNillableExpiredAt(i *int64) *TShortUrlUpdateOne {
	if i != nil {
		tuuo.SetExpiredAt(*i)
	}
	return tuuo
}

// AddExpiredAt adds i to the "expired_at" field.
func (tuuo *TShortUrlUpdateOne) AddExpiredAt(i int64) *TShortUrlUpdateOne {
	tuuo.mutation.AddExpiredAt(i)
	return tuuo
}

// Mutation returns the TShortUrlMutation object of the builder.
func (tuuo *TShortUrlUpdateOne) Mutation() *TShortUrlMutation {
	return tuuo.mutation
}

// Where appends a list predicates to the TShortUrlUpdate builder.
func (tuuo *TShortUrlUpdateOne) Where(ps ...predicate.TShortUrl) *TShortUrlUpdateOne {
	tuuo.mutation.Where(ps...)
	return tuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuuo *TShortUrlUpdateOne) Select(field string, fields ...string) *TShortUrlUpdateOne {
	tuuo.fields = append([]string{field}, fields...)
	return tuuo
}

// Save executes the query and returns the updated TShortUrl entity.
func (tuuo *TShortUrlUpdateOne) Save(ctx context.Context) (*TShortUrl, error) {
	return withHooks(ctx, tuuo.sqlSave, tuuo.mutation, tuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuuo *TShortUrlUpdateOne) SaveX(ctx context.Context) *TShortUrl {
	node, err := tuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuuo *TShortUrlUpdateOne) Exec(ctx context.Context) error {
	_, err := tuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuuo *TShortUrlUpdateOne) ExecX(ctx context.Context) {
	if err := tuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuuo *TShortUrlUpdateOne) check() error {
	if v, ok := tuuo.mutation.Origin(); ok {
		if err := tshorturl.OriginValidator(v); err != nil {
			return &ValidationError{Name: "origin", err: fmt.Errorf(`ent: validator failed for field "TShortUrl.origin": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuuo *TShortUrlUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TShortUrlUpdateOne {
	tuuo.modifiers = append(tuuo.modifiers, modifiers...)
	return tuuo
}

func (tuuo *TShortUrlUpdateOne) sqlSave(ctx context.Context) (_node *TShortUrl, err error) {
	if err := tuuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tshorturl.Table, tshorturl.Columns, sqlgraph.NewFieldSpec(tshorturl.FieldID, field.TypeInt64))
	id, ok := tuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TShortUrl.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tshorturl.FieldID)
		for _, f := range fields {
			if !tshorturl.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tshorturl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuuo.mutation.BizID(); ok {
		_spec.SetField(tshorturl.FieldBizID, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.AddedBizID(); ok {
		_spec.AddField(tshorturl.FieldBizID, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.Origin(); ok {
		_spec.SetField(tshorturl.FieldOrigin, field.TypeString, value)
	}
	if value, ok := tuuo.mutation.Visit(); ok {
		_spec.SetField(tshorturl.FieldVisit, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.AddedVisit(); ok {
		_spec.AddField(tshorturl.FieldVisit, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tshorturl.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tshorturl.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.ExpiredAt(); ok {
		_spec.SetField(tshorturl.FieldExpiredAt, field.TypeInt64, value)
	}
	if value, ok := tuuo.mutation.AddedExpiredAt(); ok {
		_spec.AddField(tshorturl.FieldExpiredAt, field.TypeInt64, value)
	}
	_spec.AddModifiers(tuuo.modifiers...)
	_node = &TShortUrl{config: tuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tshorturl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuuo.mutation.done = true
	return _node, nil
}

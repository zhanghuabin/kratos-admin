// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/position"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (pc *PositionCreate) SetCreateTime(t time.Time) *PositionCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreateTime(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PositionCreate) SetUpdateTime(t time.Time) *PositionCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PositionCreate) SetNillableUpdateTime(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetDeleteTime sets the "delete_time" field.
func (pc *PositionCreate) SetDeleteTime(t time.Time) *PositionCreate {
	pc.mutation.SetDeleteTime(t)
	return pc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (pc *PositionCreate) SetNillableDeleteTime(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetDeleteTime(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PositionCreate) SetStatus(po position.Status) *PositionCreate {
	pc.mutation.SetStatus(po)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PositionCreate) SetNillableStatus(po *position.Status) *PositionCreate {
	if po != nil {
		pc.SetStatus(*po)
	}
	return pc
}

// SetCreateBy sets the "create_by" field.
func (pc *PositionCreate) SetCreateBy(u uint32) *PositionCreate {
	pc.mutation.SetCreateBy(u)
	return pc
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreateBy(u *uint32) *PositionCreate {
	if u != nil {
		pc.SetCreateBy(*u)
	}
	return pc
}

// SetRemark sets the "remark" field.
func (pc *PositionCreate) SetRemark(s string) *PositionCreate {
	pc.mutation.SetRemark(s)
	return pc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pc *PositionCreate) SetNillableRemark(s *string) *PositionCreate {
	if s != nil {
		pc.SetRemark(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PositionCreate) SetName(s string) *PositionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PositionCreate) SetNillableName(s *string) *PositionCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetCode sets the "code" field.
func (pc *PositionCreate) SetCode(s string) *PositionCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCode(s *string) *PositionCreate {
	if s != nil {
		pc.SetCode(*s)
	}
	return pc
}

// SetParentID sets the "parent_id" field.
func (pc *PositionCreate) SetParentID(u uint32) *PositionCreate {
	pc.mutation.SetParentID(u)
	return pc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pc *PositionCreate) SetNillableParentID(u *uint32) *PositionCreate {
	if u != nil {
		pc.SetParentID(*u)
	}
	return pc
}

// SetOrderNo sets the "order_no" field.
func (pc *PositionCreate) SetOrderNo(i int32) *PositionCreate {
	pc.mutation.SetOrderNo(i)
	return pc
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (pc *PositionCreate) SetNillableOrderNo(i *int32) *PositionCreate {
	if i != nil {
		pc.SetOrderNo(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PositionCreate) SetID(u uint32) *PositionCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetParent sets the "parent" edge to the Position entity.
func (pc *PositionCreate) SetParent(p *Position) *PositionCreate {
	return pc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the Position entity by IDs.
func (pc *PositionCreate) AddChildIDs(ids ...uint32) *PositionCreate {
	pc.mutation.AddChildIDs(ids...)
	return pc
}

// AddChildren adds the "children" edges to the Position entity.
func (pc *PositionCreate) AddChildren(p ...*Position) *PositionCreate {
	ids := make([]uint32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddChildIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PositionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PositionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PositionCreate) defaults() {
	if _, ok := pc.mutation.Status(); !ok {
		v := position.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Remark(); !ok {
		v := position.DefaultRemark
		pc.mutation.SetRemark(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := position.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Code(); !ok {
		v := position.DefaultCode
		pc.mutation.SetCode(v)
	}
	if _, ok := pc.mutation.ParentID(); !ok {
		v := position.DefaultParentID
		pc.mutation.SetParentID(v)
	}
	if _, ok := pc.mutation.OrderNo(); !ok {
		v := position.DefaultOrderNo
		pc.mutation.SetOrderNo(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PositionCreate) check() error {
	if v, ok := pc.mutation.Status(); ok {
		if err := position.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Position.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Position.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Position.code"`)}
	}
	if v, ok := pc.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	if _, ok := pc.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`ent: missing required field "Position.order_no"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := position.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Position.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		_node = &Position{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(position.Table, sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(position.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := pc.mutation.DeleteTime(); ok {
		_spec.SetField(position.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := pc.mutation.CreateBy(); ok {
		_spec.SetField(position.FieldCreateBy, field.TypeUint32, value)
		_node.CreateBy = &value
	}
	if value, ok := pc.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.OrderNo(); ok {
		_spec.SetField(position.FieldOrderNo, field.TypeInt32, value)
		_node.OrderNo = value
	}
	if nodes := pc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Position.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PositionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PositionCreate) OnConflict(opts ...sql.ConflictOption) *PositionUpsertOne {
	pc.conflict = opts
	return &PositionUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Position.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PositionCreate) OnConflictColumns(columns ...string) *PositionUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PositionUpsertOne{
		create: pc,
	}
}

type (
	// PositionUpsertOne is the builder for "upsert"-ing
	//  one Position node.
	PositionUpsertOne struct {
		create *PositionCreate
	}

	// PositionUpsert is the "OnConflict" setter.
	PositionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *PositionUpsert) SetUpdateTime(v time.Time) *PositionUpsert {
	u.Set(position.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PositionUpsert) UpdateUpdateTime() *PositionUpsert {
	u.SetExcluded(position.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PositionUpsert) ClearUpdateTime() *PositionUpsert {
	u.SetNull(position.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *PositionUpsert) SetDeleteTime(v time.Time) *PositionUpsert {
	u.Set(position.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PositionUpsert) UpdateDeleteTime() *PositionUpsert {
	u.SetExcluded(position.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PositionUpsert) ClearDeleteTime() *PositionUpsert {
	u.SetNull(position.FieldDeleteTime)
	return u
}

// SetStatus sets the "status" field.
func (u *PositionUpsert) SetStatus(v position.Status) *PositionUpsert {
	u.Set(position.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PositionUpsert) UpdateStatus() *PositionUpsert {
	u.SetExcluded(position.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *PositionUpsert) ClearStatus() *PositionUpsert {
	u.SetNull(position.FieldStatus)
	return u
}

// SetCreateBy sets the "create_by" field.
func (u *PositionUpsert) SetCreateBy(v uint32) *PositionUpsert {
	u.Set(position.FieldCreateBy, v)
	return u
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *PositionUpsert) UpdateCreateBy() *PositionUpsert {
	u.SetExcluded(position.FieldCreateBy)
	return u
}

// AddCreateBy adds v to the "create_by" field.
func (u *PositionUpsert) AddCreateBy(v uint32) *PositionUpsert {
	u.Add(position.FieldCreateBy, v)
	return u
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *PositionUpsert) ClearCreateBy() *PositionUpsert {
	u.SetNull(position.FieldCreateBy)
	return u
}

// SetRemark sets the "remark" field.
func (u *PositionUpsert) SetRemark(v string) *PositionUpsert {
	u.Set(position.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *PositionUpsert) UpdateRemark() *PositionUpsert {
	u.SetExcluded(position.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *PositionUpsert) ClearRemark() *PositionUpsert {
	u.SetNull(position.FieldRemark)
	return u
}

// SetName sets the "name" field.
func (u *PositionUpsert) SetName(v string) *PositionUpsert {
	u.Set(position.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PositionUpsert) UpdateName() *PositionUpsert {
	u.SetExcluded(position.FieldName)
	return u
}

// SetCode sets the "code" field.
func (u *PositionUpsert) SetCode(v string) *PositionUpsert {
	u.Set(position.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *PositionUpsert) UpdateCode() *PositionUpsert {
	u.SetExcluded(position.FieldCode)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *PositionUpsert) SetParentID(v uint32) *PositionUpsert {
	u.Set(position.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PositionUpsert) UpdateParentID() *PositionUpsert {
	u.SetExcluded(position.FieldParentID)
	return u
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PositionUpsert) ClearParentID() *PositionUpsert {
	u.SetNull(position.FieldParentID)
	return u
}

// SetOrderNo sets the "order_no" field.
func (u *PositionUpsert) SetOrderNo(v int32) *PositionUpsert {
	u.Set(position.FieldOrderNo, v)
	return u
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *PositionUpsert) UpdateOrderNo() *PositionUpsert {
	u.SetExcluded(position.FieldOrderNo)
	return u
}

// AddOrderNo adds v to the "order_no" field.
func (u *PositionUpsert) AddOrderNo(v int32) *PositionUpsert {
	u.Add(position.FieldOrderNo, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Position.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(position.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PositionUpsertOne) UpdateNewValues() *PositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(position.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(position.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Position.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PositionUpsertOne) Ignore() *PositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PositionUpsertOne) DoNothing() *PositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PositionCreate.OnConflict
// documentation for more info.
func (u *PositionUpsertOne) Update(set func(*PositionUpsert)) *PositionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PositionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PositionUpsertOne) SetUpdateTime(v time.Time) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateUpdateTime() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PositionUpsertOne) ClearUpdateTime() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *PositionUpsertOne) SetDeleteTime(v time.Time) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateDeleteTime() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PositionUpsertOne) ClearDeleteTime() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearDeleteTime()
	})
}

// SetStatus sets the "status" field.
func (u *PositionUpsertOne) SetStatus(v position.Status) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateStatus() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *PositionUpsertOne) ClearStatus() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearStatus()
	})
}

// SetCreateBy sets the "create_by" field.
func (u *PositionUpsertOne) SetCreateBy(v uint32) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetCreateBy(v)
	})
}

// AddCreateBy adds v to the "create_by" field.
func (u *PositionUpsertOne) AddCreateBy(v uint32) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.AddCreateBy(v)
	})
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateCreateBy() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateCreateBy()
	})
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *PositionUpsertOne) ClearCreateBy() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearCreateBy()
	})
}

// SetRemark sets the "remark" field.
func (u *PositionUpsertOne) SetRemark(v string) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateRemark() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *PositionUpsertOne) ClearRemark() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearRemark()
	})
}

// SetName sets the "name" field.
func (u *PositionUpsertOne) SetName(v string) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateName() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateName()
	})
}

// SetCode sets the "code" field.
func (u *PositionUpsertOne) SetCode(v string) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateCode() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateCode()
	})
}

// SetParentID sets the "parent_id" field.
func (u *PositionUpsertOne) SetParentID(v uint32) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateParentID() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PositionUpsertOne) ClearParentID() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.ClearParentID()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *PositionUpsertOne) SetOrderNo(v int32) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *PositionUpsertOne) AddOrderNo(v int32) *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *PositionUpsertOne) UpdateOrderNo() *PositionUpsertOne {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateOrderNo()
	})
}

// Exec executes the query.
func (u *PositionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PositionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PositionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PositionUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PositionUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PositionCreateBulk is the builder for creating many Position entities in bulk.
type PositionCreateBulk struct {
	config
	err      error
	builders []*PositionCreate
	conflict []sql.ConflictOption
}

// Save creates the Position entities in the database.
func (pcb *PositionCreateBulk) Save(ctx context.Context) ([]*Position, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Position, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PositionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
					nodes[i].ID = uint32(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PositionCreateBulk) SaveX(ctx context.Context) []*Position {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PositionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PositionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Position.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PositionUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PositionCreateBulk) OnConflict(opts ...sql.ConflictOption) *PositionUpsertBulk {
	pcb.conflict = opts
	return &PositionUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Position.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PositionCreateBulk) OnConflictColumns(columns ...string) *PositionUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PositionUpsertBulk{
		create: pcb,
	}
}

// PositionUpsertBulk is the builder for "upsert"-ing
// a bulk of Position nodes.
type PositionUpsertBulk struct {
	create *PositionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Position.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(position.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PositionUpsertBulk) UpdateNewValues() *PositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(position.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(position.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Position.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PositionUpsertBulk) Ignore() *PositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PositionUpsertBulk) DoNothing() *PositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PositionCreateBulk.OnConflict
// documentation for more info.
func (u *PositionUpsertBulk) Update(set func(*PositionUpsert)) *PositionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PositionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PositionUpsertBulk) SetUpdateTime(v time.Time) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateUpdateTime() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *PositionUpsertBulk) ClearUpdateTime() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *PositionUpsertBulk) SetDeleteTime(v time.Time) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateDeleteTime() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *PositionUpsertBulk) ClearDeleteTime() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearDeleteTime()
	})
}

// SetStatus sets the "status" field.
func (u *PositionUpsertBulk) SetStatus(v position.Status) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateStatus() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *PositionUpsertBulk) ClearStatus() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearStatus()
	})
}

// SetCreateBy sets the "create_by" field.
func (u *PositionUpsertBulk) SetCreateBy(v uint32) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetCreateBy(v)
	})
}

// AddCreateBy adds v to the "create_by" field.
func (u *PositionUpsertBulk) AddCreateBy(v uint32) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.AddCreateBy(v)
	})
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateCreateBy() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateCreateBy()
	})
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *PositionUpsertBulk) ClearCreateBy() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearCreateBy()
	})
}

// SetRemark sets the "remark" field.
func (u *PositionUpsertBulk) SetRemark(v string) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateRemark() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *PositionUpsertBulk) ClearRemark() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearRemark()
	})
}

// SetName sets the "name" field.
func (u *PositionUpsertBulk) SetName(v string) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateName() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateName()
	})
}

// SetCode sets the "code" field.
func (u *PositionUpsertBulk) SetCode(v string) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateCode() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateCode()
	})
}

// SetParentID sets the "parent_id" field.
func (u *PositionUpsertBulk) SetParentID(v uint32) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateParentID() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *PositionUpsertBulk) ClearParentID() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.ClearParentID()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *PositionUpsertBulk) SetOrderNo(v int32) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *PositionUpsertBulk) AddOrderNo(v int32) *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *PositionUpsertBulk) UpdateOrderNo() *PositionUpsertBulk {
	return u.Update(func(s *PositionUpsert) {
		s.UpdateOrderNo()
	})
}

// Exec executes the query.
func (u *PositionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PositionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PositionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PositionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/organization"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (oc *OrganizationCreate) SetCreateTime(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreateTime(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *OrganizationCreate) SetUpdateTime(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdateTime(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetDeleteTime sets the "delete_time" field.
func (oc *OrganizationCreate) SetDeleteTime(t time.Time) *OrganizationCreate {
	oc.mutation.SetDeleteTime(t)
	return oc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDeleteTime(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetDeleteTime(*t)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrganizationCreate) SetStatus(o organization.Status) *OrganizationCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableStatus(o *organization.Status) *OrganizationCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetCreateBy sets the "create_by" field.
func (oc *OrganizationCreate) SetCreateBy(u uint32) *OrganizationCreate {
	oc.mutation.SetCreateBy(u)
	return oc
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreateBy(u *uint32) *OrganizationCreate {
	if u != nil {
		oc.SetCreateBy(*u)
	}
	return oc
}

// SetRemark sets the "remark" field.
func (oc *OrganizationCreate) SetRemark(s string) *OrganizationCreate {
	oc.mutation.SetRemark(s)
	return oc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableRemark(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetRemark(*s)
	}
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableName(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetName(*s)
	}
	return oc
}

// SetParentID sets the "parent_id" field.
func (oc *OrganizationCreate) SetParentID(u uint32) *OrganizationCreate {
	oc.mutation.SetParentID(u)
	return oc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableParentID(u *uint32) *OrganizationCreate {
	if u != nil {
		oc.SetParentID(*u)
	}
	return oc
}

// SetOrderNo sets the "order_no" field.
func (oc *OrganizationCreate) SetOrderNo(i int32) *OrganizationCreate {
	oc.mutation.SetOrderNo(i)
	return oc
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableOrderNo(i *int32) *OrganizationCreate {
	if i != nil {
		oc.SetOrderNo(*i)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrganizationCreate) SetID(u uint32) *OrganizationCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetParent sets the "parent" edge to the Organization entity.
func (oc *OrganizationCreate) SetParent(o *Organization) *OrganizationCreate {
	return oc.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (oc *OrganizationCreate) AddChildIDs(ids ...uint32) *OrganizationCreate {
	oc.mutation.AddChildIDs(ids...)
	return oc
}

// AddChildren adds the "children" edges to the Organization entity.
func (oc *OrganizationCreate) AddChildren(o ...*Organization) *OrganizationCreate {
	ids := make([]uint32, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddChildIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() {
	if _, ok := oc.mutation.Status(); !ok {
		v := organization.DefaultStatus
		oc.mutation.SetStatus(v)
	}
	if _, ok := oc.mutation.Remark(); !ok {
		v := organization.DefaultRemark
		oc.mutation.SetRemark(v)
	}
	if _, ok := oc.mutation.Name(); !ok {
		v := organization.DefaultName
		oc.mutation.SetName(v)
	}
	if _, ok := oc.mutation.OrderNo(); !ok {
		v := organization.DefaultOrderNo
		oc.mutation.SetOrderNo(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if v, ok := oc.mutation.Status(); ok {
		if err := organization.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Organization.status": %w`, err)}
		}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := oc.mutation.ID(); ok {
		if err := organization.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Organization.id": %w`, err)}
		}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = oc.conflict
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.SetField(organization.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.SetField(organization.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := oc.mutation.DeleteTime(); ok {
		_spec.SetField(organization.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = &value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(organization.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := oc.mutation.CreateBy(); ok {
		_spec.SetField(organization.FieldCreateBy, field.TypeUint32, value)
		_node.CreateBy = &value
	}
	if value, ok := oc.mutation.Remark(); ok {
		_spec.SetField(organization.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := oc.mutation.OrderNo(); ok {
		_spec.SetField(organization.FieldOrderNo, field.TypeInt32, value)
		_node.OrderNo = &value
	}
	if nodes := oc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUint32),
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
//	client.Organization.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrganizationUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (oc *OrganizationCreate) OnConflict(opts ...sql.ConflictOption) *OrganizationUpsertOne {
	oc.conflict = opts
	return &OrganizationUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oc *OrganizationCreate) OnConflictColumns(columns ...string) *OrganizationUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OrganizationUpsertOne{
		create: oc,
	}
}

type (
	// OrganizationUpsertOne is the builder for "upsert"-ing
	//  one Organization node.
	OrganizationUpsertOne struct {
		create *OrganizationCreate
	}

	// OrganizationUpsert is the "OnConflict" setter.
	OrganizationUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsert) SetUpdateTime(v time.Time) *OrganizationUpsert {
	u.Set(organization.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateUpdateTime() *OrganizationUpsert {
	u.SetExcluded(organization.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsert) ClearUpdateTime() *OrganizationUpsert {
	u.SetNull(organization.FieldUpdateTime)
	return u
}

// SetDeleteTime sets the "delete_time" field.
func (u *OrganizationUpsert) SetDeleteTime(v time.Time) *OrganizationUpsert {
	u.Set(organization.FieldDeleteTime, v)
	return u
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateDeleteTime() *OrganizationUpsert {
	u.SetExcluded(organization.FieldDeleteTime)
	return u
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *OrganizationUpsert) ClearDeleteTime() *OrganizationUpsert {
	u.SetNull(organization.FieldDeleteTime)
	return u
}

// SetStatus sets the "status" field.
func (u *OrganizationUpsert) SetStatus(v organization.Status) *OrganizationUpsert {
	u.Set(organization.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateStatus() *OrganizationUpsert {
	u.SetExcluded(organization.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *OrganizationUpsert) ClearStatus() *OrganizationUpsert {
	u.SetNull(organization.FieldStatus)
	return u
}

// SetCreateBy sets the "create_by" field.
func (u *OrganizationUpsert) SetCreateBy(v uint32) *OrganizationUpsert {
	u.Set(organization.FieldCreateBy, v)
	return u
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateCreateBy() *OrganizationUpsert {
	u.SetExcluded(organization.FieldCreateBy)
	return u
}

// AddCreateBy adds v to the "create_by" field.
func (u *OrganizationUpsert) AddCreateBy(v uint32) *OrganizationUpsert {
	u.Add(organization.FieldCreateBy, v)
	return u
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *OrganizationUpsert) ClearCreateBy() *OrganizationUpsert {
	u.SetNull(organization.FieldCreateBy)
	return u
}

// SetRemark sets the "remark" field.
func (u *OrganizationUpsert) SetRemark(v string) *OrganizationUpsert {
	u.Set(organization.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateRemark() *OrganizationUpsert {
	u.SetExcluded(organization.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *OrganizationUpsert) ClearRemark() *OrganizationUpsert {
	u.SetNull(organization.FieldRemark)
	return u
}

// SetName sets the "name" field.
func (u *OrganizationUpsert) SetName(v string) *OrganizationUpsert {
	u.Set(organization.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateName() *OrganizationUpsert {
	u.SetExcluded(organization.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *OrganizationUpsert) ClearName() *OrganizationUpsert {
	u.SetNull(organization.FieldName)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *OrganizationUpsert) SetParentID(v uint32) *OrganizationUpsert {
	u.Set(organization.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateParentID() *OrganizationUpsert {
	u.SetExcluded(organization.FieldParentID)
	return u
}

// ClearParentID clears the value of the "parent_id" field.
func (u *OrganizationUpsert) ClearParentID() *OrganizationUpsert {
	u.SetNull(organization.FieldParentID)
	return u
}

// SetOrderNo sets the "order_no" field.
func (u *OrganizationUpsert) SetOrderNo(v int32) *OrganizationUpsert {
	u.Set(organization.FieldOrderNo, v)
	return u
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateOrderNo() *OrganizationUpsert {
	u.SetExcluded(organization.FieldOrderNo)
	return u
}

// AddOrderNo adds v to the "order_no" field.
func (u *OrganizationUpsert) AddOrderNo(v int32) *OrganizationUpsert {
	u.Add(organization.FieldOrderNo, v)
	return u
}

// ClearOrderNo clears the value of the "order_no" field.
func (u *OrganizationUpsert) ClearOrderNo() *OrganizationUpsert {
	u.SetNull(organization.FieldOrderNo)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(organization.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrganizationUpsertOne) UpdateNewValues() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(organization.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(organization.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Organization.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrganizationUpsertOne) Ignore() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrganizationUpsertOne) DoNothing() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrganizationCreate.OnConflict
// documentation for more info.
func (u *OrganizationUpsertOne) Update(set func(*OrganizationUpsert)) *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrganizationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsertOne) SetUpdateTime(v time.Time) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateUpdateTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsertOne) ClearUpdateTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *OrganizationUpsertOne) SetDeleteTime(v time.Time) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateDeleteTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *OrganizationUpsertOne) ClearDeleteTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearDeleteTime()
	})
}

// SetStatus sets the "status" field.
func (u *OrganizationUpsertOne) SetStatus(v organization.Status) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateStatus() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *OrganizationUpsertOne) ClearStatus() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearStatus()
	})
}

// SetCreateBy sets the "create_by" field.
func (u *OrganizationUpsertOne) SetCreateBy(v uint32) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateBy(v)
	})
}

// AddCreateBy adds v to the "create_by" field.
func (u *OrganizationUpsertOne) AddCreateBy(v uint32) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.AddCreateBy(v)
	})
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateCreateBy() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateBy()
	})
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *OrganizationUpsertOne) ClearCreateBy() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearCreateBy()
	})
}

// SetRemark sets the "remark" field.
func (u *OrganizationUpsertOne) SetRemark(v string) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateRemark() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrganizationUpsertOne) ClearRemark() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearRemark()
	})
}

// SetName sets the "name" field.
func (u *OrganizationUpsertOne) SetName(v string) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateName() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *OrganizationUpsertOne) ClearName() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *OrganizationUpsertOne) SetParentID(v uint32) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateParentID() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *OrganizationUpsertOne) ClearParentID() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearParentID()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *OrganizationUpsertOne) SetOrderNo(v int32) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *OrganizationUpsertOne) AddOrderNo(v int32) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateOrderNo() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateOrderNo()
	})
}

// ClearOrderNo clears the value of the "order_no" field.
func (u *OrganizationUpsertOne) ClearOrderNo() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearOrderNo()
	})
}

// Exec executes the query.
func (u *OrganizationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrganizationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrganizationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrganizationUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrganizationUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	err      error
	builders []*OrganizationCreate
	conflict []sql.ConflictOption
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Organization.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrganizationUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ocb *OrganizationCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrganizationUpsertBulk {
	ocb.conflict = opts
	return &OrganizationUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ocb *OrganizationCreateBulk) OnConflictColumns(columns ...string) *OrganizationUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OrganizationUpsertBulk{
		create: ocb,
	}
}

// OrganizationUpsertBulk is the builder for "upsert"-ing
// a bulk of Organization nodes.
type OrganizationUpsertBulk struct {
	create *OrganizationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(organization.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrganizationUpsertBulk) UpdateNewValues() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(organization.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(organization.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrganizationUpsertBulk) Ignore() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrganizationUpsertBulk) DoNothing() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrganizationCreateBulk.OnConflict
// documentation for more info.
func (u *OrganizationUpsertBulk) Update(set func(*OrganizationUpsert)) *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrganizationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsertBulk) SetUpdateTime(v time.Time) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateUpdateTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsertBulk) ClearUpdateTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetDeleteTime sets the "delete_time" field.
func (u *OrganizationUpsertBulk) SetDeleteTime(v time.Time) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetDeleteTime(v)
	})
}

// UpdateDeleteTime sets the "delete_time" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateDeleteTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateDeleteTime()
	})
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (u *OrganizationUpsertBulk) ClearDeleteTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearDeleteTime()
	})
}

// SetStatus sets the "status" field.
func (u *OrganizationUpsertBulk) SetStatus(v organization.Status) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateStatus() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *OrganizationUpsertBulk) ClearStatus() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearStatus()
	})
}

// SetCreateBy sets the "create_by" field.
func (u *OrganizationUpsertBulk) SetCreateBy(v uint32) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateBy(v)
	})
}

// AddCreateBy adds v to the "create_by" field.
func (u *OrganizationUpsertBulk) AddCreateBy(v uint32) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.AddCreateBy(v)
	})
}

// UpdateCreateBy sets the "create_by" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateCreateBy() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateBy()
	})
}

// ClearCreateBy clears the value of the "create_by" field.
func (u *OrganizationUpsertBulk) ClearCreateBy() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearCreateBy()
	})
}

// SetRemark sets the "remark" field.
func (u *OrganizationUpsertBulk) SetRemark(v string) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateRemark() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrganizationUpsertBulk) ClearRemark() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearRemark()
	})
}

// SetName sets the "name" field.
func (u *OrganizationUpsertBulk) SetName(v string) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateName() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *OrganizationUpsertBulk) ClearName() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *OrganizationUpsertBulk) SetParentID(v uint32) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateParentID() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateParentID()
	})
}

// ClearParentID clears the value of the "parent_id" field.
func (u *OrganizationUpsertBulk) ClearParentID() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearParentID()
	})
}

// SetOrderNo sets the "order_no" field.
func (u *OrganizationUpsertBulk) SetOrderNo(v int32) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetOrderNo(v)
	})
}

// AddOrderNo adds v to the "order_no" field.
func (u *OrganizationUpsertBulk) AddOrderNo(v int32) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.AddOrderNo(v)
	})
}

// UpdateOrderNo sets the "order_no" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateOrderNo() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateOrderNo()
	})
}

// ClearOrderNo clears the value of the "order_no" field.
func (u *OrganizationUpsertBulk) ClearOrderNo() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearOrderNo()
	})
}

// Exec executes the query.
func (u *OrganizationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrganizationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrganizationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrganizationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

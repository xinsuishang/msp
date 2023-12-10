// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"msp/biz_server/oss/internal/infra/mysql/model/ent/domainmapping"
	"msp/biz_server/oss/internal/infra/mysql/model/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DomainMappingUpdate is the builder for updating DomainMapping entities.
type DomainMappingUpdate struct {
	config
	hooks    []Hook
	mutation *DomainMappingMutation
}

// Where appends a list predicates to the DomainMappingUpdate builder.
func (dmu *DomainMappingUpdate) Where(ps ...predicate.DomainMapping) *DomainMappingUpdate {
	dmu.mutation.Where(ps...)
	return dmu
}

// SetUpdatedAt sets the "updated_at" field.
func (dmu *DomainMappingUpdate) SetUpdatedAt(t time.Time) *DomainMappingUpdate {
	dmu.mutation.SetUpdatedAt(t)
	return dmu
}

// SetTenantID sets the "tenant_id" field.
func (dmu *DomainMappingUpdate) SetTenantID(i int) *DomainMappingUpdate {
	dmu.mutation.ResetTenantID()
	dmu.mutation.SetTenantID(i)
	return dmu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (dmu *DomainMappingUpdate) SetNillableTenantID(i *int) *DomainMappingUpdate {
	if i != nil {
		dmu.SetTenantID(*i)
	}
	return dmu
}

// AddTenantID adds i to the "tenant_id" field.
func (dmu *DomainMappingUpdate) AddTenantID(i int) *DomainMappingUpdate {
	dmu.mutation.AddTenantID(i)
	return dmu
}

// SetRegionID sets the "region_id" field.
func (dmu *DomainMappingUpdate) SetRegionID(s string) *DomainMappingUpdate {
	dmu.mutation.SetRegionID(s)
	return dmu
}

// SetNillableRegionID sets the "region_id" field if the given value is not nil.
func (dmu *DomainMappingUpdate) SetNillableRegionID(s *string) *DomainMappingUpdate {
	if s != nil {
		dmu.SetRegionID(*s)
	}
	return dmu
}

// SetDomain sets the "domain" field.
func (dmu *DomainMappingUpdate) SetDomain(s string) *DomainMappingUpdate {
	dmu.mutation.SetDomain(s)
	return dmu
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (dmu *DomainMappingUpdate) SetNillableDomain(s *string) *DomainMappingUpdate {
	if s != nil {
		dmu.SetDomain(*s)
	}
	return dmu
}

// SetBucketName sets the "bucket_name" field.
func (dmu *DomainMappingUpdate) SetBucketName(s string) *DomainMappingUpdate {
	dmu.mutation.SetBucketName(s)
	return dmu
}

// SetNillableBucketName sets the "bucket_name" field if the given value is not nil.
func (dmu *DomainMappingUpdate) SetNillableBucketName(s *string) *DomainMappingUpdate {
	if s != nil {
		dmu.SetBucketName(*s)
	}
	return dmu
}

// SetDesc sets the "desc" field.
func (dmu *DomainMappingUpdate) SetDesc(s string) *DomainMappingUpdate {
	dmu.mutation.SetDesc(s)
	return dmu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (dmu *DomainMappingUpdate) SetNillableDesc(s *string) *DomainMappingUpdate {
	if s != nil {
		dmu.SetDesc(*s)
	}
	return dmu
}

// Mutation returns the DomainMappingMutation object of the builder.
func (dmu *DomainMappingUpdate) Mutation() *DomainMappingMutation {
	return dmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dmu *DomainMappingUpdate) Save(ctx context.Context) (int, error) {
	dmu.defaults()
	return withHooks(ctx, dmu.sqlSave, dmu.mutation, dmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dmu *DomainMappingUpdate) SaveX(ctx context.Context) int {
	affected, err := dmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dmu *DomainMappingUpdate) Exec(ctx context.Context) error {
	_, err := dmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmu *DomainMappingUpdate) ExecX(ctx context.Context) {
	if err := dmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmu *DomainMappingUpdate) defaults() {
	if _, ok := dmu.mutation.UpdatedAt(); !ok {
		v := domainmapping.UpdateDefaultUpdatedAt()
		dmu.mutation.SetUpdatedAt(v)
	}
}

func (dmu *DomainMappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(domainmapping.Table, domainmapping.Columns, sqlgraph.NewFieldSpec(domainmapping.FieldID, field.TypeInt))
	if ps := dmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dmu.mutation.UpdatedAt(); ok {
		_spec.SetField(domainmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dmu.mutation.TenantID(); ok {
		_spec.SetField(domainmapping.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := dmu.mutation.AddedTenantID(); ok {
		_spec.AddField(domainmapping.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := dmu.mutation.RegionID(); ok {
		_spec.SetField(domainmapping.FieldRegionID, field.TypeString, value)
	}
	if value, ok := dmu.mutation.Domain(); ok {
		_spec.SetField(domainmapping.FieldDomain, field.TypeString, value)
	}
	if value, ok := dmu.mutation.BucketName(); ok {
		_spec.SetField(domainmapping.FieldBucketName, field.TypeString, value)
	}
	if value, ok := dmu.mutation.Desc(); ok {
		_spec.SetField(domainmapping.FieldDesc, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domainmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dmu.mutation.done = true
	return n, nil
}

// DomainMappingUpdateOne is the builder for updating a single DomainMapping entity.
type DomainMappingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DomainMappingMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dmuo *DomainMappingUpdateOne) SetUpdatedAt(t time.Time) *DomainMappingUpdateOne {
	dmuo.mutation.SetUpdatedAt(t)
	return dmuo
}

// SetTenantID sets the "tenant_id" field.
func (dmuo *DomainMappingUpdateOne) SetTenantID(i int) *DomainMappingUpdateOne {
	dmuo.mutation.ResetTenantID()
	dmuo.mutation.SetTenantID(i)
	return dmuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (dmuo *DomainMappingUpdateOne) SetNillableTenantID(i *int) *DomainMappingUpdateOne {
	if i != nil {
		dmuo.SetTenantID(*i)
	}
	return dmuo
}

// AddTenantID adds i to the "tenant_id" field.
func (dmuo *DomainMappingUpdateOne) AddTenantID(i int) *DomainMappingUpdateOne {
	dmuo.mutation.AddTenantID(i)
	return dmuo
}

// SetRegionID sets the "region_id" field.
func (dmuo *DomainMappingUpdateOne) SetRegionID(s string) *DomainMappingUpdateOne {
	dmuo.mutation.SetRegionID(s)
	return dmuo
}

// SetNillableRegionID sets the "region_id" field if the given value is not nil.
func (dmuo *DomainMappingUpdateOne) SetNillableRegionID(s *string) *DomainMappingUpdateOne {
	if s != nil {
		dmuo.SetRegionID(*s)
	}
	return dmuo
}

// SetDomain sets the "domain" field.
func (dmuo *DomainMappingUpdateOne) SetDomain(s string) *DomainMappingUpdateOne {
	dmuo.mutation.SetDomain(s)
	return dmuo
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (dmuo *DomainMappingUpdateOne) SetNillableDomain(s *string) *DomainMappingUpdateOne {
	if s != nil {
		dmuo.SetDomain(*s)
	}
	return dmuo
}

// SetBucketName sets the "bucket_name" field.
func (dmuo *DomainMappingUpdateOne) SetBucketName(s string) *DomainMappingUpdateOne {
	dmuo.mutation.SetBucketName(s)
	return dmuo
}

// SetNillableBucketName sets the "bucket_name" field if the given value is not nil.
func (dmuo *DomainMappingUpdateOne) SetNillableBucketName(s *string) *DomainMappingUpdateOne {
	if s != nil {
		dmuo.SetBucketName(*s)
	}
	return dmuo
}

// SetDesc sets the "desc" field.
func (dmuo *DomainMappingUpdateOne) SetDesc(s string) *DomainMappingUpdateOne {
	dmuo.mutation.SetDesc(s)
	return dmuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (dmuo *DomainMappingUpdateOne) SetNillableDesc(s *string) *DomainMappingUpdateOne {
	if s != nil {
		dmuo.SetDesc(*s)
	}
	return dmuo
}

// Mutation returns the DomainMappingMutation object of the builder.
func (dmuo *DomainMappingUpdateOne) Mutation() *DomainMappingMutation {
	return dmuo.mutation
}

// Where appends a list predicates to the DomainMappingUpdate builder.
func (dmuo *DomainMappingUpdateOne) Where(ps ...predicate.DomainMapping) *DomainMappingUpdateOne {
	dmuo.mutation.Where(ps...)
	return dmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dmuo *DomainMappingUpdateOne) Select(field string, fields ...string) *DomainMappingUpdateOne {
	dmuo.fields = append([]string{field}, fields...)
	return dmuo
}

// Save executes the query and returns the updated DomainMapping entity.
func (dmuo *DomainMappingUpdateOne) Save(ctx context.Context) (*DomainMapping, error) {
	dmuo.defaults()
	return withHooks(ctx, dmuo.sqlSave, dmuo.mutation, dmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dmuo *DomainMappingUpdateOne) SaveX(ctx context.Context) *DomainMapping {
	node, err := dmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dmuo *DomainMappingUpdateOne) Exec(ctx context.Context) error {
	_, err := dmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmuo *DomainMappingUpdateOne) ExecX(ctx context.Context) {
	if err := dmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmuo *DomainMappingUpdateOne) defaults() {
	if _, ok := dmuo.mutation.UpdatedAt(); !ok {
		v := domainmapping.UpdateDefaultUpdatedAt()
		dmuo.mutation.SetUpdatedAt(v)
	}
}

func (dmuo *DomainMappingUpdateOne) sqlSave(ctx context.Context) (_node *DomainMapping, err error) {
	_spec := sqlgraph.NewUpdateSpec(domainmapping.Table, domainmapping.Columns, sqlgraph.NewFieldSpec(domainmapping.FieldID, field.TypeInt))
	id, ok := dmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DomainMapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, domainmapping.FieldID)
		for _, f := range fields {
			if !domainmapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != domainmapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(domainmapping.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dmuo.mutation.TenantID(); ok {
		_spec.SetField(domainmapping.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := dmuo.mutation.AddedTenantID(); ok {
		_spec.AddField(domainmapping.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := dmuo.mutation.RegionID(); ok {
		_spec.SetField(domainmapping.FieldRegionID, field.TypeString, value)
	}
	if value, ok := dmuo.mutation.Domain(); ok {
		_spec.SetField(domainmapping.FieldDomain, field.TypeString, value)
	}
	if value, ok := dmuo.mutation.BucketName(); ok {
		_spec.SetField(domainmapping.FieldBucketName, field.TypeString, value)
	}
	if value, ok := dmuo.mutation.Desc(); ok {
		_spec.SetField(domainmapping.FieldDesc, field.TypeString, value)
	}
	_node = &DomainMapping{config: dmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domainmapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dmuo.mutation.done = true
	return _node, nil
}

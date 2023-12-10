// Code generated by ent, DO NOT EDIT.

package domainmapping

import (
	"msp/biz_server/oss/internal/infra/mysql/model/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldUpdatedAt, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldTenantID, v))
}

// RegionID applies equality check predicate on the "region_id" field. It's identical to RegionIDEQ.
func RegionID(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldRegionID, v))
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldDomain, v))
}

// BucketName applies equality check predicate on the "bucket_name" field. It's identical to BucketNameEQ.
func BucketName(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldBucketName, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldDesc, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldUpdatedAt, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v int) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldTenantID, v))
}

// RegionIDEQ applies the EQ predicate on the "region_id" field.
func RegionIDEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldRegionID, v))
}

// RegionIDNEQ applies the NEQ predicate on the "region_id" field.
func RegionIDNEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldRegionID, v))
}

// RegionIDIn applies the In predicate on the "region_id" field.
func RegionIDIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldRegionID, vs...))
}

// RegionIDNotIn applies the NotIn predicate on the "region_id" field.
func RegionIDNotIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldRegionID, vs...))
}

// RegionIDGT applies the GT predicate on the "region_id" field.
func RegionIDGT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldRegionID, v))
}

// RegionIDGTE applies the GTE predicate on the "region_id" field.
func RegionIDGTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldRegionID, v))
}

// RegionIDLT applies the LT predicate on the "region_id" field.
func RegionIDLT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldRegionID, v))
}

// RegionIDLTE applies the LTE predicate on the "region_id" field.
func RegionIDLTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldRegionID, v))
}

// RegionIDContains applies the Contains predicate on the "region_id" field.
func RegionIDContains(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContains(FieldRegionID, v))
}

// RegionIDHasPrefix applies the HasPrefix predicate on the "region_id" field.
func RegionIDHasPrefix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasPrefix(FieldRegionID, v))
}

// RegionIDHasSuffix applies the HasSuffix predicate on the "region_id" field.
func RegionIDHasSuffix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasSuffix(FieldRegionID, v))
}

// RegionIDEqualFold applies the EqualFold predicate on the "region_id" field.
func RegionIDEqualFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEqualFold(FieldRegionID, v))
}

// RegionIDContainsFold applies the ContainsFold predicate on the "region_id" field.
func RegionIDContainsFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContainsFold(FieldRegionID, v))
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldDomain, v))
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldDomain, v))
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldDomain, vs...))
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldDomain, vs...))
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldDomain, v))
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldDomain, v))
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldDomain, v))
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldDomain, v))
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContains(FieldDomain, v))
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasPrefix(FieldDomain, v))
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasSuffix(FieldDomain, v))
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEqualFold(FieldDomain, v))
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContainsFold(FieldDomain, v))
}

// BucketNameEQ applies the EQ predicate on the "bucket_name" field.
func BucketNameEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldBucketName, v))
}

// BucketNameNEQ applies the NEQ predicate on the "bucket_name" field.
func BucketNameNEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldBucketName, v))
}

// BucketNameIn applies the In predicate on the "bucket_name" field.
func BucketNameIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldBucketName, vs...))
}

// BucketNameNotIn applies the NotIn predicate on the "bucket_name" field.
func BucketNameNotIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldBucketName, vs...))
}

// BucketNameGT applies the GT predicate on the "bucket_name" field.
func BucketNameGT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldBucketName, v))
}

// BucketNameGTE applies the GTE predicate on the "bucket_name" field.
func BucketNameGTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldBucketName, v))
}

// BucketNameLT applies the LT predicate on the "bucket_name" field.
func BucketNameLT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldBucketName, v))
}

// BucketNameLTE applies the LTE predicate on the "bucket_name" field.
func BucketNameLTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldBucketName, v))
}

// BucketNameContains applies the Contains predicate on the "bucket_name" field.
func BucketNameContains(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContains(FieldBucketName, v))
}

// BucketNameHasPrefix applies the HasPrefix predicate on the "bucket_name" field.
func BucketNameHasPrefix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasPrefix(FieldBucketName, v))
}

// BucketNameHasSuffix applies the HasSuffix predicate on the "bucket_name" field.
func BucketNameHasSuffix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasSuffix(FieldBucketName, v))
}

// BucketNameEqualFold applies the EqualFold predicate on the "bucket_name" field.
func BucketNameEqualFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEqualFold(FieldBucketName, v))
}

// BucketNameContainsFold applies the ContainsFold predicate on the "bucket_name" field.
func BucketNameContainsFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContainsFold(FieldBucketName, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.DomainMapping {
	return predicate.DomainMapping(sql.FieldContainsFold(FieldDesc, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DomainMapping) predicate.DomainMapping {
	return predicate.DomainMapping(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DomainMapping) predicate.DomainMapping {
	return predicate.DomainMapping(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DomainMapping) predicate.DomainMapping {
	return predicate.DomainMapping(sql.NotPredicates(p))
}

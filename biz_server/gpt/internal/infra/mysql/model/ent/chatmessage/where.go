// Code generated by ent, DO NOT EDIT.

package chatmessage

import (
	"msp/biz_server/gpt/internal/infra/mysql/model/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldChatID, v))
}

// RequestID applies equality check predicate on the "request_id" field. It's identical to RequestIDEQ.
func RequestID(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldRequestID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldText, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldVersion, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldUpdatedAt, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int32) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldChatID, v))
}

// RequestIDEQ applies the EQ predicate on the "request_id" field.
func RequestIDEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldRequestID, v))
}

// RequestIDNEQ applies the NEQ predicate on the "request_id" field.
func RequestIDNEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldRequestID, v))
}

// RequestIDIn applies the In predicate on the "request_id" field.
func RequestIDIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldRequestID, vs...))
}

// RequestIDNotIn applies the NotIn predicate on the "request_id" field.
func RequestIDNotIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldRequestID, vs...))
}

// RequestIDGT applies the GT predicate on the "request_id" field.
func RequestIDGT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldRequestID, v))
}

// RequestIDGTE applies the GTE predicate on the "request_id" field.
func RequestIDGTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldRequestID, v))
}

// RequestIDLT applies the LT predicate on the "request_id" field.
func RequestIDLT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldRequestID, v))
}

// RequestIDLTE applies the LTE predicate on the "request_id" field.
func RequestIDLTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldRequestID, v))
}

// RequestIDContains applies the Contains predicate on the "request_id" field.
func RequestIDContains(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContains(FieldRequestID, v))
}

// RequestIDHasPrefix applies the HasPrefix predicate on the "request_id" field.
func RequestIDHasPrefix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasPrefix(FieldRequestID, v))
}

// RequestIDHasSuffix applies the HasSuffix predicate on the "request_id" field.
func RequestIDHasSuffix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasSuffix(FieldRequestID, v))
}

// RequestIDEqualFold applies the EqualFold predicate on the "request_id" field.
func RequestIDEqualFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldRequestID, v))
}

// RequestIDContainsFold applies the ContainsFold predicate on the "request_id" field.
func RequestIDContainsFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldRequestID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldText, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int8) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldVersion, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.NotPredicates(p))
}

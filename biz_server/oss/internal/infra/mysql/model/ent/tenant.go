// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"msp/biz_server/oss/internal/infra/mysql/model/ent/tenant"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Tenant is the model entity for the Tenant schema.
type Tenant struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Region holds the value of the "region" field.
	Region bool `json:"region,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// AccessKey holds the value of the "access_key" field.
	AccessKey string `json:"access_key,omitempty"`
	// SecretKey holds the value of the "secret_key" field.
	SecretKey string `json:"secret_key,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc         string `json:"desc,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tenant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tenant.FieldRegion:
			values[i] = new(sql.NullBool)
		case tenant.FieldID:
			values[i] = new(sql.NullInt64)
		case tenant.FieldName, tenant.FieldType, tenant.FieldAccessKey, tenant.FieldSecretKey, tenant.FieldDesc:
			values[i] = new(sql.NullString)
		case tenant.FieldCreatedAt, tenant.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tenant fields.
func (t *Tenant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tenant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tenant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tenant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case tenant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tenant.FieldRegion:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				t.Region = value.Bool
			}
		case tenant.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = value.String
			}
		case tenant.FieldAccessKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_key", values[i])
			} else if value.Valid {
				t.AccessKey = value.String
			}
		case tenant.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				t.SecretKey = value.String
			}
		case tenant.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				t.Desc = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tenant.
// This includes values selected through modifiers, order, etc.
func (t *Tenant) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Tenant.
// Note that you need to call Tenant.Unwrap() before calling this method if this Tenant
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tenant) Update() *TenantUpdateOne {
	return NewTenantClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tenant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tenant) Unwrap() *Tenant {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tenant is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tenant) String() string {
	var builder strings.Builder
	builder.WriteString("Tenant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(fmt.Sprintf("%v", t.Region))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(t.Type)
	builder.WriteString(", ")
	builder.WriteString("access_key=")
	builder.WriteString(t.AccessKey)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(t.SecretKey)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(t.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// Tenants is a parsable slice of Tenant.
type Tenants []*Tenant

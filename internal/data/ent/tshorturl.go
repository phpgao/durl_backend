// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
)

// TShortUrl is the model entity for the TShortUrl schema.
type TShortUrl struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// refer to leaf id
	BizID int64 `json:"biz_id,omitempty"`
	// current max id
	Origin string `json:"origin,omitempty"`
	// short url
	Short int64 `json:"short,omitempty"`
	// Visit holds the value of the "visit" field.
	Visit int64 `json:"visit,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt    int64 `json:"expired_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TShortUrl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tshorturl.FieldID, tshorturl.FieldBizID, tshorturl.FieldShort, tshorturl.FieldVisit, tshorturl.FieldCreatedAt, tshorturl.FieldUpdatedAt, tshorturl.FieldExpiredAt:
			values[i] = new(sql.NullInt64)
		case tshorturl.FieldOrigin:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TShortUrl fields.
func (tu *TShortUrl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tshorturl.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tu.ID = int64(value.Int64)
		case tshorturl.FieldBizID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field biz_id", values[i])
			} else if value.Valid {
				tu.BizID = value.Int64
			}
		case tshorturl.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				tu.Origin = value.String
			}
		case tshorturl.FieldShort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				tu.Short = value.Int64
			}
		case tshorturl.FieldVisit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field visit", values[i])
			} else if value.Valid {
				tu.Visit = value.Int64
			}
		case tshorturl.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tu.CreatedAt = value.Int64
			}
		case tshorturl.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tu.UpdatedAt = value.Int64
			}
		case tshorturl.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				tu.ExpiredAt = value.Int64
			}
		default:
			tu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TShortUrl.
// This includes values selected through modifiers, order, etc.
func (tu *TShortUrl) Value(name string) (ent.Value, error) {
	return tu.selectValues.Get(name)
}

// Update returns a builder for updating this TShortUrl.
// Note that you need to call TShortUrl.Unwrap() before calling this method if this TShortUrl
// was returned from a transaction, and the transaction was committed or rolled back.
func (tu *TShortUrl) Update() *TShortUrlUpdateOne {
	return NewTShortUrlClient(tu.config).UpdateOne(tu)
}

// Unwrap unwraps the TShortUrl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tu *TShortUrl) Unwrap() *TShortUrl {
	_tx, ok := tu.config.driver.(*txDriver)
	if !ok {
		panic("ent: TShortUrl is not a transactional entity")
	}
	tu.config.driver = _tx.drv
	return tu
}

// String implements the fmt.Stringer.
func (tu *TShortUrl) String() string {
	var builder strings.Builder
	builder.WriteString("TShortUrl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tu.ID))
	builder.WriteString("biz_id=")
	builder.WriteString(fmt.Sprintf("%v", tu.BizID))
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(tu.Origin)
	builder.WriteString(", ")
	builder.WriteString("short=")
	builder.WriteString(fmt.Sprintf("%v", tu.Short))
	builder.WriteString(", ")
	builder.WriteString("visit=")
	builder.WriteString(fmt.Sprintf("%v", tu.Visit))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("expired_at=")
	builder.WriteString(fmt.Sprintf("%v", tu.ExpiredAt))
	builder.WriteByte(')')
	return builder.String()
}

// TShortUrls is a parsable slice of TShortUrl.
type TShortUrls []*TShortUrl
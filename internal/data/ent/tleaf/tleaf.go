// Code generated by ent, DO NOT EDIT.

package tleaf

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tleaf type in the database.
	Label = "tleaf"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBizTag holds the string denoting the biz_tag field in the database.
	FieldBizTag = "biz_tag"
	// FieldMaxID holds the string denoting the max_id field in the database.
	FieldMaxID = "max_id"
	// FieldStep holds the string denoting the step field in the database.
	FieldStep = "step"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the tleaf in the database.
	Table = "leaf"
)

// Columns holds all SQL columns for tleaf fields.
var Columns = []string{
	FieldID,
	FieldBizTag,
	FieldMaxID,
	FieldStep,
	FieldDesc,
	FieldVersion,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// BizTagValidator is a validator for the "biz_tag" field. It is called by the builders before save.
	BizTagValidator func(string) error
	// DefaultMaxID holds the default value on creation for the "max_id" field.
	DefaultMaxID int64
	// DefaultStep holds the default value on creation for the "step" field.
	DefaultStep int64
	// DescValidator is a validator for the "desc" field. It is called by the builders before save.
	DescValidator func(string) error
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int32
)

// OrderOption defines the ordering options for the TLeaf queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBizTag orders the results by the biz_tag field.
func ByBizTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBizTag, opts...).ToFunc()
}

// ByMaxID orders the results by the max_id field.
func ByMaxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxID, opts...).ToFunc()
}

// ByStep orders the results by the step field.
func ByStep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStep, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

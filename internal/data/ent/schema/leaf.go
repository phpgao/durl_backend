package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TLeaf struct {
	ent.Schema
}

func (TLeaf) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "leaf"},
		entsql.WithComments(true),
	}
}

func (TLeaf) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("primary key"),
		field.String("biz_tag").NotEmpty().Unique().Comment("for different biz"),
		field.Int64("max_id").Default(1000).Comment("current max id"),
		field.Int64("step").Default(1000).Comment("nums per batch"),
		field.String("desc").MaxLen(255),
		field.Int32("version").Default(0).Comment("version control"),
		field.Int64("created_at").Immutable(),
		field.Int64("updated_at"),
	}
}

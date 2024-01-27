package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TShortUrl struct {
	ent.Schema
}

func (TShortUrl) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "short_url"},
		entsql.WithComments(true),
	}
}

func (TShortUrl) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("primary key"),
		field.Int64("biz_id").Comment("refer to leaf id"),
		field.String("origin").NotEmpty().MinLen(10).Comment("current max id"),
		field.Int64("short").Unique().Immutable().Comment("short url"),
		field.Int64("visit").Default(0),
		field.Int64("created_at").Immutable(),
		field.Int64("updated_at"),
		field.Int64("expired_at").Default(0),
	}
}

func (TShortUrl) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("short", "biz_id").
			Unique(),
	}
}

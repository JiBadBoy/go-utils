package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/JiBadBoy/go-utils/sonyflake"
)

type SnowflakeId struct {
	mixin.Schema
}

func (SnowflakeId) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Comment("id").
			DefaultFunc(sonyflake.GenerateSonyflake).
			Positive().
			Immutable().
			StructTag(`json:"id,omitempty"`).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
			}),
	}
}

// Indexes of the SnowflakeId.
func (SnowflakeId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

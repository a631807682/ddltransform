package schema

import (
	"strings"
)

type Field struct {
	DBName          string
	DBType          string
	ModelType       string
	PrimaryKey      bool
	AutoIncrement   bool
	HasDefaultValue bool
	DefaultValue    string
	NotNull         bool
	Unique          bool
	UniqueKeyName   string
	Comment         string
	GoType          FieldGoType
}

func (field *Field) GetTagString() string {
	attrs := make([]string, 2, 10)
	attrs[0] = "column:" + field.DBName
	attrs[1] = "type:" + field.DBType

	if field.PrimaryKey {
		attrs = append(attrs, "primaryKey")
	}

	if field.AutoIncrement {
		attrs = append(attrs, "autoIncrement")
	}

	if field.HasDefaultValue {
		attrs = append(attrs, "default:"+field.DefaultValue)
	}

	if field.NotNull {
		attrs = append(attrs, "NOT NULL")
	}

	if field.Unique {
		attrs = append(attrs, "uniqueIndex:"+field.UniqueKeyName)
	}

	if field.Comment != "" {
		attrs = append(attrs, "comment:"+field.Comment)
	}

	return strings.Join(attrs, ";")
}

type FieldGoType string

const (
	Bool   FieldGoType = "bool"
	Int    FieldGoType = "int"
	Uint   FieldGoType = "uint"
	Float  FieldGoType = "float"
	String FieldGoType = "string"
	Time   FieldGoType = "time"
)

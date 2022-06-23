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
}

func (filed *Field) GetTagString() string {
	attrs := make([]string, 2, 10)
	attrs[0] = "column:" + filed.DBName
	attrs[1] = "type:" + filed.DBType

	if filed.PrimaryKey {
		attrs = append(attrs, "primaryKey")
	}

	if filed.AutoIncrement {
		attrs = append(attrs, "autoIncrement")
	}

	if filed.HasDefaultValue {
		attrs = append(attrs, "default:"+filed.DefaultValue)
	}

	if filed.NotNull {
		attrs = append(attrs, "NOT NULL")
	}

	if filed.Unique {
		attrs = append(attrs, "uniqueIndex:"+filed.UniqueKeyName)
	}

	if filed.Comment != "" {
		attrs = append(attrs, "comment:"+filed.Comment)
	}

	return strings.Join(attrs, ";")
}

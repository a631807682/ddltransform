package transformer

import (
	"strings"

	"github.com/a631807682/ddltransform/schema"
	"github.com/a631807682/ddltransform/utils"

	j "github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
)

// GormTransformer gorm transformer implement
type GormTransformer struct{}

func (*GormTransformer) Name() string {
	return "gorm"
}

// Transform implement transform ddl info to code
func (*GormTransformer) Transform(table string, fields []schema.Field) (modeCode string, err error) {
	tableName := utils.ToFormatName(inflection.Singular(table))

	fCodes := make([]j.Code, 0)
	for _, field := range fields {
		tags := make(map[string]string, 1)
		tags["gorm"] = getTagString(field)
		// field := j.Id(utils.ToFormatName(field.DBName)).String().Tag(tags)
		fCode := j.Id(utils.ToFormatName(field.DBName))

		switch field.GoType {
		case schema.Bool:
			fCode.Bool()
		case schema.Int:
			fCode.Int64()
		case schema.Uint:
			fCode.Uint64()
		case schema.Float:
			fCode.Float64()
		case schema.String:
			fCode.String()
		case schema.Time:
			fCode.Qual("time", "Time")
		}

		fCode.Tag(tags)
		fCodes = append(fCodes, fCode)
	}

	c := j.Type().Id(tableName).Struct(
		fCodes...,
	)

	return c.GoString(), nil
}

func getTagString(field schema.Field) string {
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

	if field.UniqueIndex {
		attrs = append(attrs, "uniqueIndex:"+field.UniqueIndexName)
	}

	if field.Comment != "" {
		attrs = append(attrs, "comment:"+field.Comment)
	}

	return strings.Join(attrs, ";")
}

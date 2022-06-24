package transformer

import (
	"github.com/a631807682/ddltransform/schema"
	"github.com/a631807682/ddltransform/utils"

	j "github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
)

type GormTransformer struct{}

func (*GormTransformer) Name() string {
	return "gorm"
}

func (*GormTransformer) Transform(table string, fields []schema.Field) (modeCode string, err error) {
	tableName := utils.ToFormatName(inflection.Singular(table))

	fCodes := make([]j.Code, 0)
	for _, field := range fields {
		tags := make(map[string]string, 1)
		tags["gorm"] = field.GetTagString()
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

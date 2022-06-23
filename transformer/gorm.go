package transformer

import (
	"ddltransform/schema"
	"ddltransform/utils"
	"fmt"

	j "github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
)

type GormTransformer struct{}

func (*GormTransformer) Transform(table string, fileds []schema.Field) (modeCode string, err error) {
	tableName := utils.ToFormatName(inflection.Singular(table))

	fCodes := make([]j.Code, 0)
	for _, field := range fileds {
		tags := make(map[string]string, 1)
		tags["gorm"] = field.GetTagString()
		filed := j.Id(utils.ToFormatName(field.DBName)).String().Tag(tags)

		fCodes = append(fCodes, filed)
	}

	c := j.Type().Id(tableName).Struct(
		fCodes...,
	)
	fmt.Println(tableName)
	fmt.Printf("%#v", c.GoString())

	return c.GoString(), nil
}

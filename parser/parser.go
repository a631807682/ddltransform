package parser

import "ddltransform/schema"

type Parser interface {
	Name() string
	Parse(ddl string) (table string, fileds []schema.Field, err error)
}

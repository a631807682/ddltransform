package parser

import "github.com/a631807682/ddltransform/schema"

type Parser interface {
	Name() string
	Parse(ddl string) (table string, fields []schema.Field, err error)
}

package parser

import "github.com/a631807682/ddltransform/schema"

type Parser interface {
	// Name parser name
	Name() string
	// Parse parse ddl to table and column info
	Parse(ddl string) (table string, fields []schema.Field, err error)
}

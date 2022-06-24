package transformer

import "github.com/a631807682/ddltransform/schema"

type Transformer interface {
	// Name transformer name
	Name() string
	// Transform transform ddl info to code
	Transform(table string, fields []schema.Field) (code string, err error)
}

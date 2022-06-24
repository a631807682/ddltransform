package transformer

import "github.com/a631807682/ddltransform/schema"

type Transformer interface {
	Name() string
	Transform(table string, fields []schema.Field) (modeCode string, err error)
}

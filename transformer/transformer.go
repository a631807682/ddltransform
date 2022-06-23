package transformer

import "ddltransform/schema"

type Transformer interface {
	Name() string
	Transform(table string, fileds []schema.Field) (modeCode string, err error)
}

package main

import (
	"github.com/a631807682/ddltransform"
)

// use gorm transformer
func trans2gorm(ddl string) (code string, err error) {
	return ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:      ddltransform.Mysql,
		TransformerType: ddltransform.Gorm,
	})
}

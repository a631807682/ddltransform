package ddltransform

import (
	"fmt"

	"github.com/a631807682/ddltransform/parser"
	"github.com/a631807682/ddltransform/transformer"
)

type ParserType uint
type TransformerType uint

const (
	// support parser type
	Mysql  ParserType = 1
	Sqlite ParserType = 2
	// support transformer type
	Gorm TransformerType = 1
)

var parserMap = map[ParserType]parser.Parser{
	Mysql:  &parser.MysqlParser{},
	Sqlite: &parser.MysqlParser{},
}

var transformerMap = map[TransformerType]transformer.Transformer{
	Gorm: &transformer.GormTransformer{},
}

func type2parser(pt ParserType) parser.Parser {
	return parserMap[pt]
}

func type2Transformer(tt TransformerType) transformer.Transformer {
	return transformerMap[tt]
}

type Config struct {
	// ParserType support parser type
	ParserType ParserType
	// TransformerType support transformer type
	TransformerType TransformerType
	// Parser customize parser
	Parser parser.Parser
	// Transformer customize transformer
	Transformer transformer.Transformer
}

func Transform(ddl string, config Config) (code string, err error) {
	var p parser.Parser
	var t transformer.Transformer

	if config.Parser != nil { //customize parser
		p = config.Parser
	} else {
		p = type2parser(config.ParserType)
	}

	if p == nil {
		err = fmt.Errorf("parser missing. config:%+v", config)
		return
	}

	if config.Transformer != nil { //customize transformer
		t = config.Transformer
	} else {
		t = type2Transformer(config.TransformerType)
	}

	if t == nil {
		err = fmt.Errorf("transformer missing. config:%+v", config)
		return
	}

	table, fields, err := p.Parse(ddl)
	if err != nil {
		err = fmt.Errorf("parse failed. parser:%s err:%v", p.Name(), err)
		return
	}

	code, err = t.Transform(table, fields)
	if err != nil {
		err = fmt.Errorf("transform failed. transformer:%s err:%v", t.Name(), err)
		return
	}
	return
}

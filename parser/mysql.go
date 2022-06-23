package parser

import (
	"ddltransform/schema"
	"fmt"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	tdriver "github.com/pingcap/tidb/parser/test_driver"
)

const (
	TypeUnspecified byte = 0
	TypeTiny        byte = 1 // TINYINT
	TypeShort       byte = 2 // SMALLINT
	TypeLong        byte = 3 // INT
	TypeFloat       byte = 4
	TypeDouble      byte = 5
	TypeNull        byte = 6
	TypeTimestamp   byte = 7
	TypeLonglong    byte = 8 // BIGINT
	TypeInt24       byte = 9 // MEDIUMINT
	TypeDate        byte = 10
	/* TypeDuration original name was TypeTime, renamed to TypeDuration to resolve the conflict with Go type Time.*/
	TypeDuration byte = 11
	TypeDatetime byte = 12
	TypeYear     byte = 13
	TypeNewDate  byte = 14
	TypeVarchar  byte = 15
	TypeBit      byte = 16

	TypeJSON       byte = 0xf5
	TypeNewDecimal byte = 0xf6
	TypeEnum       byte = 0xf7
	TypeSet        byte = 0xf8
	TypeTinyBlob   byte = 0xf9
	TypeMediumBlob byte = 0xfa
	TypeLongBlob   byte = 0xfb
	TypeBlob       byte = 0xfc
	TypeVarString  byte = 0xfd
	TypeString     byte = 0xfe
	TypeGeometry   byte = 0xff
)

const (
	UnsignedFlag uint = 1 << 5 /* Field is unsigned */
)

type createTableVisitor struct {
	table  string
	fileds []schema.Field
}

func (v *createTableVisitor) Enter(in ast.Node) (ast.Node, bool) {
	if ctStmt, ok := in.(*ast.CreateTableStmt); ok {
		v.table = ctStmt.Table.Name.String()

		// primary key
		primaryKeyMaps := make(map[string]interface{})
		uniqueMaps := make(map[string]string)
		for _, c := range ctStmt.Constraints {
			switch c.Tp {
			case ast.ConstraintPrimaryKey:
				for _, k := range c.Keys {
					primaryKeyMaps[k.Column.Name.L] = nil
				}
			case ast.ConstraintUniq, ast.ConstraintUniqIndex,
				ast.ConstraintUniqKey:
				for _, k := range c.Keys {
					uniqueMaps[k.Column.Name.L] = c.Name
				}
			}
		}

		v.fileds = make([]schema.Field, 0, 10)
		for _, c := range ctStmt.Cols {
			filed := schema.Field{
				DBName: c.Name.Name.L,
				DBType: c.Tp.String(),
			}

			// database type to go type
			switch c.Tp.GetType() {
			case TypeString, TypeBlob, TypeMediumBlob, TypeLongBlob,
				TypeVarString, TypeVarchar, TypeTinyBlob:
				filed.GoType = schema.String
			case TypeDate, TypeDatetime, TypeDuration, TypeTimestamp:
				filed.GoType = schema.Time
			case TypeFloat, TypeDouble:
				filed.GoType = schema.Float
			case TypeNewDecimal:
				if c.Tp.GetDecimal() > 0 {
					filed.GoType = schema.Float
				} else {
					filed.GoType = schema.Int
				}
			case TypeTiny:
				if c.Tp.GetFlen() == 1 {
					filed.GoType = schema.Bool
				} else {
					filed.GoType = schema.Int
				}
			case TypeInt24, TypeShort, TypeLong, TypeLonglong:
				if (c.Tp.GetFlag() & UnsignedFlag) > 0 { // UNSIGNED
					filed.GoType = schema.Uint
				} else {
					filed.GoType = schema.Int
				}
			default:
				// use string to receive unhandle database type
				filed.GoType = schema.String
			}

			if _, ok := primaryKeyMaps[c.Name.Name.L]; ok {
				filed.PrimaryKey = true
			}

			filed.UniqueKeyName, filed.Unique = uniqueMaps[c.Name.Name.L]

			// filed options
			for _, opt := range c.Options {
				switch opt.Tp {
				case ast.ColumnOptionAutoIncrement:
					filed.AutoIncrement = true
				case ast.ColumnOptionNotNull:
					filed.NotNull = true
				case ast.ColumnOptionDefaultValue:
					if ve, ok := opt.Expr.(*tdriver.ValueExpr); ok {
						if ve.Datum.Kind() != tdriver.KindNull {
							filed.HasDefaultValue = true
							filed.DefaultValue = ve.Datum.GetString()
						}
					}
				case ast.ColumnOptionComment:
					if ve, ok := opt.Expr.(*tdriver.ValueExpr); ok {
						if ve.Datum.Kind() == tdriver.KindString {
							filed.Comment = ve.Datum.GetString()
						}
					}
				}
			}

			v.fileds = append(v.fileds, filed)
		}
	}

	return in, true
}

func (v *createTableVisitor) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}

type MysqlParser struct{}

func (*MysqlParser) Parse(ddl string) (table string, fileds []schema.Field, err error) {
	astNode, err := parse(ddl)
	if err != nil {
		return
	}

	v := &createTableVisitor{}
	_, ok := (*astNode).Accept(v)
	if !ok {
		err = fmt.Errorf("parse failed: not create table stmt ddl:\n%s", ddl)
		return
	}

	table = v.table
	fileds = v.fileds
	return
}

func (*MysqlParser) Name() string {
	return "mysql"
}

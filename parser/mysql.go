package parser

import (
	"ddltransform/schema"
	"fmt"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	tdriver "github.com/pingcap/tidb/parser/test_driver"
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

			if _, ok := primaryKeyMaps[c.Name.Name.L]; ok {
				filed.PrimaryKey = true
			}

			filed.UniqueKeyName, filed.Unique = uniqueMaps[c.Name.Name.L]

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

			fmt.Println(filed.DBName, len(c.Options))
			fmt.Printf("%+v\n", filed)
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

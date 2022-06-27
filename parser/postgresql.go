package parser

import (
	"fmt"
	"strings"

	"github.com/a631807682/ddltransform/schema"
	"github.com/auxten/postgresql-parser/pkg/sql/parser"
	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	"github.com/auxten/postgresql-parser/pkg/sql/types"
	"github.com/auxten/postgresql-parser/pkg/walk"
)

// PostgresqlParser postgresql parser implement
// Note: postgresql ddl not support index and comment and collate
// https://stackoverflow.com/questions/6239657/can-you-create-an-index-in-the-create-table-definition
// It's using github.com/auxten/postgresql-parser package which is not support `COLLATE`
type PostgresqlParser struct{}

// Parse implement parse ddl info
func (*PostgresqlParser) Parse(ddl string) (table string, fields []schema.Field, err error) {
	stmts, err := parser.Parse(ddl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fieldsMap := make(map[string]indexField, 0)
	w := &walk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			// create table
			if ct, ok := node.(*tree.CreateTable); ok {
				table = ct.Table.Table()

				var fieldIdx int
				for _, d := range ct.Defs {
					if columnDef, ok := d.(*tree.ColumnTableDef); ok {
						field := schema.Field{
							DBName:          columnDef.Name.String(),
							DBType:          strings.ToLower(columnDef.Type.SQLString()),
							PrimaryKey:      columnDef.PrimaryKey.IsPrimaryKey,
							AutoIncrement:   false,
							HasDefaultValue: columnDef.HasDefaultExpr(),
							NotNull:         columnDef.Nullable.Nullability == tree.NotNull,
							UniqueIndex:     false, // not support
							UniqueIndexName: "",    // not support
							Comment:         "",    // not support
							GoType:          transColumn2GoType(columnDef.Type),
						}

						if field.HasDefaultValue {
							val := columnDef.DefaultExpr.Expr.String()
							field.DefaultValue = val
							if strings.Contains(val, "nextval") {
								field.AutoIncrement = true
							}
						}

						fieldsMap[field.DBName] = indexField{
							index: fieldIdx,
							field: &field,
						}
						fieldIdx++
					} else if uniqueDef, ok := d.(*tree.UniqueConstraintTableDef); ok {
						for _, col := range uniqueDef.Columns {

							if idxField, ok := fieldsMap[col.Column.String()]; ok {
								if uniqueDef.PrimaryKey {
									idxField.field.PrimaryKey = true
								} else {
									idxField.field.Unique = true
									idxField.field.UniqueName = uniqueDef.Name.String()
								}
							}
						}
					}
				}
				return true
			}
			return false
		},
	}

	_, _ = w.Walk(stmts, nil)

	fields = make([]schema.Field, len(fieldsMap))
	for _, idxField := range fieldsMap {
		fields[idxField.index] = *idxField.field
	}

	return
}

func (*PostgresqlParser) Name() string {
	return "postgresql"
}

type indexField struct {
	index int
	field *schema.Field
}

func transColumn2GoType(typ *types.T) schema.FieldGoType {
	switch typ.Family() {
	case types.BoolFamily:
		return schema.Bool
	case types.IntFamily:
		return schema.Int
	case types.FloatFamily:
		return schema.Float
	case types.DecimalFamily:
		if typ.Width() == 0 {
			return schema.Int
		} else {
			return schema.Float
		}
	case types.DateFamily, types.TimestampFamily,
		types.IntervalFamily, types.TimestampTZFamily,
		types.TimeFamily, types.TimeTZFamily:
		return schema.Time
	case types.StringFamily, types.BytesFamily,
		types.CollatedStringFamily, types.UuidFamily,
		types.ArrayFamily, types.INetFamily, types.JsonFamily:
		return schema.String
		// case types.OidFamily, types.UnknownFamily,
		// 	types.TupleFamily, types.BitFamily,
		// 	types.AnyFamily:
	}

	// default go type
	return schema.String
}

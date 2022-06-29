package tests_test

import (
	"testing"

	"github.com/a631807682/ddltransform/schema"
	"github.com/a631807682/ddltransform/transformer"

	"github.com/stretchr/testify/assert"
)

type transTestCase struct {
	table   string
	fields  []schema.Field
	success bool
	result  string
}

func TestGormTransformer(t *testing.T) {
	gt := &transformer.GormTransformer{}
	tcs := []transTestCase{{
		table: `test_data`,
		fields: []schema.Field{{
			DBName:        "id",
			DBType:        "bigint(20) UNSIGNED",
			GoType:        schema.Uint,
			PrimaryKey:    true,
			AutoIncrement: true,
			NotNull:       true,
		}, {
			DBName:  "create_at",
			DBType:  "datetime",
			GoType:  schema.Time,
			NotNull: true,
		}, {
			DBName:  "deleted",
			DBType:  "tinyint(1)",
			GoType:  schema.Bool,
			NotNull: true,
		}, {
			DBName:          "version",
			DBType:          "bigint(20)",
			GoType:          schema.Int,
			HasDefaultValue: true,
			DefaultValue:    "10",
			Comment:         "version info",
			UniqueIndex:     true,
			UniqueIndexName: "uk_app_version",
		}, {
			DBName:          "address",
			DBType:          "varchar(255)",
			GoType:          schema.String,
			NotNull:         true,
			HasDefaultValue: true,
			DefaultValue:    "china",
		}, {
			DBName: "amount",
			DBType: "decimal(19,2)",
			GoType: schema.Float,
		}, {
			DBName:          "wx_mp_app_id",
			DBType:          "varchar(32)",
			GoType:          schema.String,
			UniqueIndex:     true,
			UniqueIndexName: "uk_app_version",
		}, {
			DBName: "contacts",
			DBType: "varchar(50)",
			GoType: schema.String,
		}},
		success: true,
		result:  "type TestDatum struct {\n\tID        uint64    `gorm:\"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;NOT NULL\"`\n\tCreateAt  time.Time `gorm:\"column:create_at;type:datetime;NOT NULL\"`\n\tDeleted   bool      `gorm:\"column:deleted;type:tinyint(1);NOT NULL\"`\n\tVersion   int64     `gorm:\"column:version;type:bigint(20);default:10;uniqueIndex:uk_app_version;comment:version info\"`\n\tAddress   string    `gorm:\"column:address;type:varchar(255);default:china;NOT NULL\"`\n\tAmount    float64   `gorm:\"column:amount;type:decimal(19,2)\"`\n\tWxMpAppID string    `gorm:\"column:wx_mp_app_id;type:varchar(32);uniqueIndex:uk_app_version\"`\n\tContacts  string    `gorm:\"column:contacts;type:varchar(50)\"`\n}",
	}}

	for _, c := range tcs {
		modelCode, err := gt.Transform(c.table, c.fields)
		assert.Equal(t, nil, err)
		assert.Equal(t, c.result, modelCode)
	}
}

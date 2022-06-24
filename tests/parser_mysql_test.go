package tests_test

import (
	"testing"

	"github.com/a631807682/ddltransform/parser"
	"github.com/a631807682/ddltransform/schema"

	"github.com/stretchr/testify/assert"
)

type parseResult struct {
	table  string
	fields []schema.Field
}

type parseTestCase struct {
	ddl      string
	success  bool
	parseRes parseResult
}

func TestMysqlParse(t *testing.T) {
	p := &parser.MysqlParser{}
	tcs := []parseTestCase{{
		ddl: `
		CREATE TABLE test_data (
			id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
			create_at datetime NOT NULL,
			deleted tinyint(1) NOT NULL,
			version bigint(20) DEFAULT '10' COMMENT 'version info',
			address varchar(255) NOT NULL DEFAULT 'china',
			amount decimal(19,2) DEFAULT NULL,
			wx_mp_app_id varchar(32) DEFAULT NULL,
			contacts varchar(50) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY uk_app_version (wx_mp_app_id, version)
		) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI ROW_FORMAT=COMPACT COMMENT='' CHECKSUM=0 DELAY_KEY_WRITE=0;
		`,
		success: true,
		parseRes: parseResult{
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
				Unique:          true,
				UniqueKeyName:   "uk_app_version",
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
				DBName:        "wx_mp_app_id",
				DBType:        "varchar(32)",
				GoType:        schema.String,
				Unique:        true,
				UniqueKeyName: "uk_app_version",
			}, {
				DBName: "contacts",
				DBType: "varchar(50)",
				GoType: schema.String,
			}},
		},
	}}
	testParseCases(t, p, tcs)
}

func testParseCases(t *testing.T, p parser.Parser, tcs []parseTestCase) {
	t.Helper()
	for _, c := range tcs {
		table, fields, err := p.Parse(c.ddl)
		if (c.success && err != nil) || (!c.success && err == nil) {
			t.Errorf("success flag and parse status not equal expect status:%v got err:%v", c.success, err)
		}

		assert.Equal(t, c.parseRes.table, table)
		assert.Equal(t, len(c.parseRes.fields), len(fields))
		for i := 0; i < len(c.parseRes.fields); i++ {
			assert.Equalf(t, c.parseRes.fields[i], fields[i], "index:%d name:%s", i, c.parseRes.fields[i].DBName)
		}
	}
}

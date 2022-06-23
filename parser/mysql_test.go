package parser

import (
	"ddltransform/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseResult struct {
	table  string
	fileds []schema.Field
}

type testCase struct {
	ddl     string
	success bool
	res     testCaseResult
}

func TestMysqlParse(t *testing.T) {
	p := &MysqlParser{}
	tcs := []testCase{{
		ddl: `
		CREATE TABLE test_data (
			id bigint(20) NOT NULL AUTO_INCREMENT,
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
		res: testCaseResult{
			table: `test_data`,
			fileds: []schema.Field{{
				DBName:        "id",
				DBType:        "bigint(20)",
				PrimaryKey:    true,
				AutoIncrement: true,
				NotNull:       true,
			}, {
				DBName:  "create_at",
				DBType:  "datetime",
				NotNull: true,
			}, {
				DBName:  "deleted",
				DBType:  "tinyint(1)",
				NotNull: true,
			}, {
				DBName:          "version",
				DBType:          "bigint(20)",
				HasDefaultValue: true,
				DefaultValue:    "10",
				Comment:         "version info",
				Unique:          true,
				UniqueKeyName:   "uk_app_version",
			}, {
				DBName:          "address",
				DBType:          "varchar(255)",
				NotNull:         true,
				HasDefaultValue: true,
				DefaultValue:    "china",
			}, {
				DBName: "amount",
				DBType: "decimal(19,2)",
			}, {
				DBName:        "wx_mp_app_id",
				DBType:        "varchar(32)",
				Unique:        true,
				UniqueKeyName: "uk_app_version",
			}, {
				DBName: "contacts",
				DBType: "varchar(50)",
			}},
		},
	}}

	for _, c := range tcs {
		table, fields, err := p.Parse(c.ddl)
		if (c.success && err != nil) || (!c.success && err == nil) {
			t.Errorf("success flag and parse status not equal expect status:%v got err:%v", c.success, err)
		}

		assert.Equal(t, c.res.table, table)
		assert.Equal(t, len(c.res.fileds), len(fields))
		for i := 0; i < len(c.res.fileds); i++ {
			assert.Equalf(t, c.res.fileds[i], fields[i], "index:%d name:%s", i, c.res.fileds[i].DBName)
		}
	}
}

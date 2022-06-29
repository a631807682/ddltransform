package tests_test

import (
	"testing"

	"github.com/a631807682/ddltransform"
	"github.com/a631807682/ddltransform/schema"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTransform(t *testing.T) {
	ddl := `
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
		`
	expectCode := "type TestDatum struct {\n\tID        uint64    `gorm:\"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;NOT NULL\"`\n\tCreateAt  time.Time `gorm:\"column:create_at;type:datetime;NOT NULL\"`\n\tDeleted   bool      `gorm:\"column:deleted;type:tinyint(1);NOT NULL\"`\n\tVersion   int64     `gorm:\"column:version;type:bigint(20);default:10;uniqueIndex:uk_app_version;comment:version info\"`\n\tAddress   string    `gorm:\"column:address;type:varchar(255);default:china;NOT NULL\"`\n\tAmount    float64   `gorm:\"column:amount;type:decimal(19,2)\"`\n\tWxMpAppID string    `gorm:\"column:wx_mp_app_id;type:varchar(32);uniqueIndex:uk_app_version\"`\n\tContacts  string    `gorm:\"column:contacts;type:varchar(50)\"`\n}"

	code, err := ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:      ddltransform.Mysql,
		TransformerType: ddltransform.Gorm,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, expectCode, code)
}

type emptyTransformer struct {
}

func (*emptyTransformer) Name() string {
	return "empty_transformer"
}

func (*emptyTransformer) Transform(_ string, _ []schema.Field) (string, error) {
	return "empty", nil
}
func TestCustomizeTransform(t *testing.T) {
	ddl := `
	CREATE TABLE blacklists  (
		id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
		PRIMARY KEY (id) USING BTREE
	) ENGINE = InnoDB AUTO_INCREMENT = 1;
	`
	code, err := ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:  ddltransform.Mysql,
		Transformer: &emptyTransformer{},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "empty", code)
}

func TestError(t *testing.T) {
	var err error
	_, err = ddltransform.Transform("", ddltransform.Config{
		ParserType: ddltransform.Mysql,
	})
	assert.NotEqual(t, nil, err)

	_, err = ddltransform.Transform("", ddltransform.Config{
		TransformerType: ddltransform.Gorm,
	})
	assert.NotEqual(t, nil, err)
}

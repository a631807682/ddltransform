package tests_test

import (
	"testing"

	"github.com/a631807682/ddltransform"
	"github.com/a631807682/ddltransform/schema"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTransform(t *testing.T) {
	ddl := `
	CREATE TABLE blacklists  (
		id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
		email varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
		type varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
		created_at datetime NOT NULL,
		updated_at datetime NOT NULL,
		PRIMARY KEY (id) USING BTREE
	) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = DYNAMIC;
	`
	expectCode := "type Blacklist struct {\n\tID        uint64    `gorm:\"column:id;type:int(10) UNSIGNED;primaryKey;autoIncrement;NOT NULL\"`\n\tEmail     string    `gorm:\"column:email;type:varchar(50) CHARACTER SET utf8;NOT NULL\"`\n\tType      string    `gorm:\"column:type;type:varchar(20) CHARACTER SET utf8;NOT NULL\"`\n\tCreatedAt time.Time `gorm:\"column:created_at;type:datetime;NOT NULL\"`\n\tUpdatedAt time.Time `gorm:\"column:updated_at;type:datetime;NOT NULL\"`\n}"

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

func (*emptyTransformer) Transform(table string, fields []schema.Field) (string, error) {
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

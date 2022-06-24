package main

import (
	"fmt"
	"strings"

	"github.com/a631807682/ddltransform"
	"github.com/a631807682/ddltransform/schema"
)

const ddl = `		
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

func main() {
	defaultCase()
	customizeCase()
}

// package support parser and transformer
func defaultCase() {
	code, err := ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:      ddltransform.Mysql,
		TransformerType: ddltransform.Gorm,
	})
	if err != nil {
		fmt.Printf("transform err:%v", err)
		return
	}

	fmt.Println("defaultCase:\n", code)
	// type TestDatum struct {
	//     ID        uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;NOT NULL"`
	//     CreateAt  time.Time `gorm:"column:create_at;type:datetime;NOT NULL"`
	//     Deleted   bool      `gorm:"column:deleted;type:tinyint(1);NOT NULL"`
	//     Version   int64     `gorm:"column:version;type:bigint(20);default:10;uniqueIndex:uk_app_version;comment:version info"`
	//     Address   string    `gorm:"column:address;type:varchar(255);default:china;NOT NULL"`
	//     Amount    float64   `gorm:"column:amount;type:decimal(19,2)"`
	//     WxMpAppID string    `gorm:"column:wx_mp_app_id;type:varchar(32);uniqueIndex:uk_app_version"`
	//     Contacts  string    `gorm:"column:contacts;type:varchar(50)"`
	// }
}

type selectTransformer struct {
}

func (*selectTransformer) Name() string {
	return "select_transfomer"
}

func (*selectTransformer) Transform(table string, fields []schema.Field) (modeCode string, err error) {
	layout := "SELECT %s FROM %s"
	cols := make([]string, len(fields))
	for i, f := range fields {
		cols[i] = f.DBName
	}
	modeCode = fmt.Sprintf(layout, strings.Join(cols, ","), table)
	return
}

// customize parser or transformer
func customizeCase() {
	code, err := ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:  ddltransform.Mysql,
		Transformer: &selectTransformer{},
	})
	if err != nil {
		fmt.Printf("transform err:%v", err)
		return
	}

	fmt.Println("customizeCase:\n", code)
	// SELECT id,create_at,deleted,version,address,amount,wx_mp_app_id,contacts FROM test_data
}

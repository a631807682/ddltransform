package main

import (
	"fmt"
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
	code, _ := trans2gorm(ddl)
	fmt.Printf("trans2gorm:\n%s\n", code)
	// SELECT id,create_at,deleted,version,address,amount,wx_mp_app_id,contacts FROM test_data
	code, _ = trans2select(ddl)
	fmt.Printf("trans2select:\n%s\n", code)
	// 	INSERT INTO test_data (id,create_at,deleted,version,address,amount,wx_mp_app_id,contacts) VALUES
	// (0,'2032-09-04 03:36:50',0,95,'DaFpLSjFbc',0.69,'atyyiNKARe','mBTvKSJfjz'),
	// (0,'2016-04-30 11:32:31',1,66,'XoEFfRsWxP',0.03,'KJyiXJrscc','aLbtZsyMGe'),
	// (0,'1999-03-20 11:10:21',0,28,'LDnJObCsNV',0.54,'tNswYNsGRu','uDtRzQMDQi'),
	// (0,'2060-08-02 00:34:11',1,58,'lgTeMaPEZQ',0.98,'ssVmaozFZB','YCOhgHOvgS'),
	// (0,'2062-06-02 17:52:17',0,47,'leQYhYzRyW',0.75,'sbOJiFQGZs','eycJPJHYNu'),
	// (0,'2027-12-30 06:55:20',0,47,'JjPjzpfRFE',0.29,'nwTKSmVoiG','fNjJhhjUVR'),
	// (0,'2026-05-03 07:09:18',1,87,'gmotaFetHs',0.75,'LOpbUOpEdK','uSqfgqVMkP'),
	// (0,'1977-01-15 17:49:08',1,88,'bZRjxAwnwe',0.15,'updOMeRVja','YVkURUpiFv'),
	// (0,'2067-05-07 18:46:56',1,90,'krBEmfdzdc',0.36,'RzLNTXYeUC','IZRgBmyArK'),
	// (0,'2009-09-23 21:37:29',0,15,'EkXBAkjQZL',0.83,'WKsXbGyRAO','CtzkjkZIva')
	code, _ = trans2insert(ddl)
	fmt.Printf("trans2insert:\n%s\n", code)
}

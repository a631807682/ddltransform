# ddltransform
Parse ddl and transform to gorm model

## Desc
Generate the orm model through parse sql to reduce the dependence on the environment

## Usage
1. Use `parser` and `transformer` package to generate model code
```go
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

code, err := ddltransform.Transform(ddl, ddltransform.Config{
	Parser:      ddltransform.Mysql,
	Transformer: ddltransform.Gorm,
})
...

fmt.Print(code)
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
```

2. Use command-line to generate model code (WIP)


## More Examples
See full list of [examples](./examples/)
package transformer

// import (
// 	"ddltransform/parser"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type testCase struct {
// 	src string
// 	ok  bool
// 	dst string
// }

// func TestTransform(t *testing.T) {
// 	tcs := []testCase{{`
// 	CREATE TABLE test_data (
// 		id bigint(20) NOT NULL AUTO_INCREMENT,
// 		create_at datetime NOT NULL,
// 		deleted tinyint(1) NOT NULL,
// 		update_at datetime NOT NULL,
// 		version bigint(20) DEFAULT '10' COMMENT 'version info',
// 		address varchar(255) NOT NULL DEFAULT 'china',
// 		amount decimal(19,2) DEFAULT NULL,
// 		charge_id varchar(32) DEFAULT NULL,
// 		paid_amount decimal(19,2) DEFAULT NULL,
// 		transaction_no varchar(64) DEFAULT NULL,
// 		wx_mp_app_id varchar(32) DEFAULT NULL,
// 		contacts varchar(50) DEFAULT NULL,
// 		deliver_fee decimal(19,2) DEFAULT NULL,
// 		deliver_info varchar(255) DEFAULT NULL,
// 		deliver_time varchar(255) DEFAULT NULL,
// 		description varchar(255) DEFAULT NULL,
// 		invoice varchar(255) DEFAULT NULL,
// 		order_from int(11) DEFAULT NULL,
// 		order_state int(11) NOT NULL,
// 		packing_fee decimal(19,2) DEFAULT NULL,
// 		payment_time datetime DEFAULT NULL,
// 		payment_type int(11) DEFAULT NULL,
// 		phone varchar(50) NOT NULL,
// 		store_employee_id bigint(20) DEFAULT NULL,
// 		store_id bigint(20) NOT NULL,
// 		user_id bigint(20) NOT NULL,
// 		payment_mode int(11) NOT NULL,
// 		current_latitude double NOT NULL,
// 		current_longitude double NOT NULL,
// 		address_latitude double NOT NULL,
// 		address_longitude double NOT NULL,
// 		PRIMARY KEY (id),
// 		UNIQUE KEY uk_phone (phone),
// 		UNIQUE KEY uk_app_version (wx_mp_app_id, version),
// 		CONSTRAINT food_order_ibfk_1 FOREIGN KEY (user_id) REFERENCES waimaiqa.user (id),
// 		CONSTRAINT food_order_ibfk_2 FOREIGN KEY (store_id) REFERENCES waimaiqa.store (id),
// 		CONSTRAINT food_order_ibfk_3 FOREIGN KEY (store_employee_id) REFERENCES waimaiqa.store_employee (id),
// 		UNIQUE FK_UNIQUE_charge_id USING BTREE (charge_id) comment '',
// 		INDEX FK_eqst2x1xisn3o0wbrlahnnqq8 USING BTREE (store_employee_id) comment '',
// 		INDEX FK_8jcmec4kb03f4dod0uqwm54o9 USING BTREE (store_id) comment '',
// 		INDEX FK_a3t0m9apja9jmrn60uab30pqd USING BTREE (user_id) comment ''
// 		) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI ROW_FORMAT=COMPACT COMMENT='' CHECKSUM=0 DELAY_KEY_WRITE=0;
// 	`, true, ``,
// 	}}

// 	for _, tc := range tcs {
// 		p := &parser.MysqlParser{}

// 		dst, err := Transform("mysql", tc.src)
// 		if (tc.ok && err != nil) || (!tc.ok && err == nil) {
// 			t.Fatalf("transform status not equal tc.ok:%v err:%v", tc.ok, err)
// 		}
// 		assert.Equal(t, dst, tc.dst)
// 	}
// }

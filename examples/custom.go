package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/a631807682/ddltransform"
	"github.com/a631807682/ddltransform/schema"
)

// customize transformer to generate select sql for select all columns
func trans2select(ddl string) (code string, err error) {
	return ddltransform.Transform(ddl, ddltransform.Config{
		ParserType:  ddltransform.Mysql,
		Transformer: &selectTransformer{},
	})
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

// customize transformer to generate insert sql for fill in data
func trans2insert(ddl string) (code string, err error) {
	return ddltransform.Transform(ddl, ddltransform.Config{
		ParserType: ddltransform.Mysql,
		Transformer: &fillInTransformer{
			Size: 10,
		},
	})
}

type fillInTransformer struct {
	Size int
}

func (*fillInTransformer) Name() string {
	return "fill_in_transfomer"
}

func (t *fillInTransformer) Transform(table string, fields []schema.Field) (modeCode string, err error) {
	layout := "INSERT INTO %s (%s) VALUES %s"
	colNames := make([]string, 0, len(fields))
	colVals := make([][]string, t.Size)
	for _, f := range fields {
		colNames = append(colNames, f.DBName)

		for i := 0; i < t.Size; i++ {
			if f.PrimaryKey {
				colVals[i] = append(colVals[i], "0")
				continue
			}
			var strVal string
			switch f.GoType {
			case schema.Bool:
				strVal = strconv.Itoa(rand.Intn(2))
			case schema.Int:
				strVal = strconv.Itoa(int(rand.Intn(100)))
			case schema.Float:
				strVal = fmt.Sprintf("%.2f", rand.Float32())
			case schema.String:
				strVal = "'" + randStringBytes(10) + "'"
			case schema.Uint:
				strVal = strconv.Itoa(int(rand.Intn(100)))
			case schema.Time:
				strVal = "'" + randate().Format("2006-01-02 15:04:05") + "'"
			}
			colVals[i] = append(colVals[i], strVal)
		}

	}
	strColNames := strings.Join(colNames, ",")
	strColVals := ""
	for i, vals := range colVals {
		strColVals += "\n(" + strings.Join(vals, ",") + ")"
		if i < len(colVals)-1 {
			strColVals += ","
		}
	}
	modeCode = fmt.Sprintf(layout, table, strColNames, strColVals)
	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

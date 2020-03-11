package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lhlyu/got/v2/db/core"
	"log"
	"strings"
)

const (
	query_tables = `SELECT TABLE_NAME,TABLE_COMMENT
					FROM INFORMATION_SCHEMA.TABLES 
					WHERE TABLE_SCHEMA=? AND (ENGINE='MyISAM' OR ENGINE = 'InnoDB' OR ENGINE = 'TokuDB')`

	query_columns = `SELECT TABLE_NAME, COLUMN_NAME,ORDINAL_POSITION,COLUMN_DEFAULT,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,CHARACTER_OCTET_LENGTH,NUMERIC_PRECISION,NUMERIC_SCALE,DATETIME_PRECISION,COLUMN_TYPE,COLUMN_KEY,EXTRA,COLUMN_COMMENT,IS_NULLABLE
					FROM INFORMATION_SCHEMA.COLUMNS
					WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`

	query_indexs = `SELECT INDEX_NAME, NON_UNIQUE, COLUMN_NAME FROM INFORMATION_SCHEMA.STATISTICS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`
)

var dict = map[string]string{
	"bigint":             "int64",
	"bigint unsigned":    "uint64",
	"binary":             "[]byte",
	"bit":                "bool",
	"blob":               "[]byte",
	"bool":               "bool",
	"boolean":            "bool",
	"char":               "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"decimal":            "float64",
	"double":             "float64",
	"enum":               "string",
	"float":              "float64",
	"int":                "int",
	"int unsigned":       "uint",
	"longblob":           "[]byte",
	"longtext":           "string",
	"mediumblob":         "[]byte",
	"mediumint":          "int",
	"mediumint unsigned": "uint",
	"mediumtext":         "string",
	"numeric":            "float64",
	"real":               "float64",
	"set":                "string",
	"smallint":           "int",
	"smallint unsigned":  "uint",
	"text":               "string",
	"time":               "time.Time",
	"timestamp":          "time.Time",
	"tinyblob":           "[]byte",
	"tinyint":            "int",
	"tinyint unsigned":   "uint",
	"varbinary":          "[]byte",
	"varchar":            "string",
	"year":               "time.Time",
}

type Mysql struct {
	db   *sql.DB
	cfg  *core.Config
	dict map[string]string
}

func (d *Mysql) Connect(cfg *core.Config) error {
	DB, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	//验证连接
	if err := DB.Ping(); err != nil {
		return err
	}
	d.db = DB
	d.cfg = cfg
	d.dict = dict
	return nil
}

func (d *Mysql) GetTables() []*core.Table {
	rows, err := d.db.Query(query_tables, d.cfg.Database)
	if err != nil {
		log.Fatalln(err)
	}
	var tables []*core.Table
	for rows.Next() {
		table := core.NewTable()
		rows.Scan(&table.Name, &table.Comment)
		tables = append(tables, table)
	}
	for _, v := range tables {
		cols := d.GetColumns(v.Name)
		v.AddColumns(cols)
	}
	return tables
}

func (d *Mysql) GetColumns(tableName string) []*core.Column {
	rows, err := d.db.Query(query_columns, d.cfg.Database, tableName)
	if err != nil {
		log.Fatalln(err)
	}
	var cols []*core.Column
	for rows.Next() {

		var tableName, name, def, dataType, columnType, columnKey, extra, comment, nullable sql.NullString
		var seq, maxLength, octetLength, numPrecision, numScale, datetimePrecision sql.NullInt64
		rows.Scan(
			&tableName,
			&name,
			&seq,
			&def,
			&dataType,
			&maxLength,
			&octetLength,
			&numPrecision,
			&numScale,
			&datetimePrecision,
			&columnType,
			&columnKey,
			&extra,
			&comment,
			&nullable)
		col := &core.Column{
			TableName:         tableName.String,
			Name:              name.String,
			Sort:              int(seq.Int64),
			Default:           def.String,
			DataType:          dataType.String,
			MaxLength:         int(maxLength.Int64),
			OctetLength:       int(octetLength.Int64),
			NumPrecision:      int(numPrecision.Int64),
			NumScale:          int(numScale.Int64),
			DatetimePrecision: int(datetimePrecision.Int64),
			ColumnType:        columnType.String,
			ColumnKey:         columnKey.String,
			Extra:             extra.String,
			Comment:           comment.String,
		}
		if nullable.String != "NO" {
			col.IsNull = true
		}
		if col.NumPrecision > 0 {
			col.IsNumber = true
			if strings.Contains(col.ColumnType, "unsigned") {
				col.IsUnsigned = true
			}
		}
		cols = append(cols, col)
	}
	return cols
}

func (*Mysql) GetIndexs(tableName string) []*core.Index {
	return nil
}

func (s *Mysql) GetDict() map[string]string {
	return s.dict
}

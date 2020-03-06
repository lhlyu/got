package db

import (
	"github.com/lhlyu/got/db/mysql"
	"github.com/xormplus/xorm"
	"strings"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Sche     string
	Database string
}

// 表属性
type Table struct {
	Name    string
	Comment string
	Columns []*Column
}

// 字段属性
type Column struct {
	Name    string
	Kind    string
	Comment string
	Pk      bool
	Attr    map[string]string
}

type DB interface {
	Connect(DbConfig) *xorm.Engine
	ToTables(*xorm.Engine) []*Table
}

func NewDb(driverName string) DB {
	switch strings.ToUpper(driverName) {
	case "mysql":
		return mysql.Mysql{}
	}
	return nil
}

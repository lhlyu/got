package db

import (
	"github.com/lhlyu/got/db/core"
	"github.com/lhlyu/got/db/mysql"
	"strings"
)

type DB interface {
	Connect(cfg *core.Config)
	GetTables() []*core.Table
	GetColumns(tableName string) []*core.Column
	GetIndexs(tableName string) []*core.Index
	GetDict() map[string]string
}

func NewDb(driverName string) DB {
	switch strings.ToLower(driverName) {
	case "mysql":
		return &mysql.Mysql{}
	}
	return nil
}

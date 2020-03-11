package db

import (
	"github.com/lhlyu/got/v2/db/core"
	"github.com/lhlyu/got/v2/db/mysql"
	"strings"
)

type DB interface {
	Connect(cfg *core.Config) error
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

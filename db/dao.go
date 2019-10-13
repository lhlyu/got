package db

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Dao interface {
	SetDB(dc *DbConf)          // 设置数据源
	QueryTables() []*TableInfo // 查询所有表
}

type DaoBase struct {
	db *sqlx.DB // 数据源
	dc *DbConf  // 配置
	driverName string  // 数据库类型
}

func NewDao(dbType string) Dao {
	switch dbType {
	case "mysql", "MYSQL":
		return new(mysqlDao)
	}
	log.Fatal("new dao is err,param is illegal")
	return nil
}

type DbConf struct {
	User            string `json:"user"`            // 用户
	Password        string `json:"password"`        // 密码
	Host            string `json:"host"`            // 主机地址
	Port            string `json:"port"`            // 端口
	Database        string `json:"database"`        // 数据库
	Schema          string `json:"schema"`          // 多租户
	Tabel           string `json:"table"`           // 表名
	Charset         string `json:"charset"`         // 默认 UTF8
	ParseTime       bool   `json:"parseTime"`       // 默认 true
	MaxOpenConns    int    `json:"maxOpenConns"`    // 最大连接数 默认 1
	MaxIdleConns    int    `json:"maxIdleConns"`    // 初始化连接数 默认 1
	ConnMaxLifetime int    `json:"connMaxLifetime"` // 存活时间 默认30s
}

func NewDbConf(host, port, user, pass, database, schema, table string) *DbConf {
	return &DbConf{
		User:            user,
		Password:        pass,
		Host:            host,
		Port:            port,
		Database:        database,
		Schema:          schema,
		Tabel:           table,
		ParseTime:       true,
		MaxIdleConns:    1,
		MaxOpenConns:    1,
		ConnMaxLifetime: 30,
		Charset:         "utf8",
	}
}

type TableInfo struct {
	TabelName     string `db:"TABLE_NAME"`
	ColumnName    string `db:"COLUMN_NAME"`
	ColumnType      string `db:"COLUMN_TYPE"`
	ColumnComment string `db:"COLUMN_COMMENT"`
}

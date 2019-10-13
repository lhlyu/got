package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type mysqlDao struct {
	DaoBase
}


func (d *mysqlDao) SetDB(dc *DbConf) {
	conn := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t"
	path := fmt.Sprintf(conn, dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Database,
		dc.Charset,
		dc.ParseTime)
	db, err := sqlx.Connect("mysql", path)
	if err != nil{
		log.Fatal("mysql init db is err = ",err.Error())
		return
	}
	db.SetMaxOpenConns(dc.MaxOpenConns)
	db.SetMaxIdleConns(dc.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(dc.ConnMaxLifetime))
	d.db = db
	d.dc = dc
	d.driverName = "mysql"
}

func (d *mysqlDao) QueryTables() []*TableInfo {
	sql := "SELECT TABLE_NAME,COLUMN_NAME,COLUMN_TYPE,COLUMN_COMMENT FROM information_schema.COLUMNS WHERE table_schema = ? %s ORDER BY TABLE_NAME;"
	params := []interface{}{d.dc.Database}
	s := ""
	if d.dc.Tabel != ""{
		s = "and table_name = ?"
		params = append(params,d.dc.Tabel)
	}
	sql = fmt.Sprintf(sql,s)
	var tf []*TableInfo
	err := d.db.Select(&tf,sql,params...)
	if err != nil{
		log.Panicln("db select is err = ",err.Error())
	}
	return tf
}


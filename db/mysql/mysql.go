package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lhlyu/got/db"
	"github.com/xormplus/xorm"
	"log"
)

type Mysql struct {
}

func (Mysql) Connect(c db.DbConfig) *xorm.Engine {
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database)
	eg, err := xorm.NewMySQL(xorm.MYSQL_DRIVER, path)
	if err != nil {
		log.Fatalln(err)
	}
	return eg
}

func (Mysql) ToTables(eg *xorm.Engine) []*db.Table {
	tabs, err := eg.DBMetas()
	if err != nil {
		log.Fatalln(err)
	}
	var tt []*db.Table
	for _, tab := range tabs {
		t := &db.Table{
			Name:    tab.Name,
			Comment: tab.Comment,
		}
		var cc []*db.Column
		for _, col := range tab.Columns() {
			c := &db.Column{
				Name:    col.Name,
				Comment: col.Comment,
				Pk:      col.IsPrimaryKey,
			}
			cc = append(cc, c)
		}
		t.Columns = cc
		tt = append(tt, t)
	}
	return tt
}

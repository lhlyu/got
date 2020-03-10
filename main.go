package main

import (
	"github.com/lhlyu/got/db"
	"github.com/lhlyu/got/db/core"
)

var cfg = &core.Config{
	Host:     "gz-cdb-dso5f8qx.sql.tencentcdb.com",
	Port:     62177,
	User:     "deve",
	Password: "deve1234",
	Database: "hos",
}

func main() {
	d := db.NewDb("mysql")

	d.Connect(cfg)
	d.ToStruct(d.GetTables()...)
}

package g

func getTypeMap(driverName string) map[string]string{
	switch driverName {
	case "mysql","MYSQL":
		return mysql
	}
	return nil
}

var mysql = map[string]string{
	"bigint":"int64",
	"bigint unsigned":"uint64",
	"binary":"[]byte",
	"bit":"bool",
	"blob":"[]byte",
	"bool":"bool",
	"boolean":"bool",
	"char":"string",
	"date":"time.Time",
	"datetime":"time.Time",
	"decimal":"float64",
	"double":"float64",
	"enum":"string",
	"float":"float64",
	"int":"int",
	"int unsigned":"uint",
	"longblob":"[]byte",
	"longtext":"string",
	"mediumblob":"[]byte",
	"mediumint":"int",
	"mediumint unsigned":"uint",
	"mediumtext":"string",
	"numeric":"float64",
	"real":"float64",
	"set":"string",
	"smallint":"int",
	"smallint unsigned":"uint",
	"text":"string",
	"time":"time.Time",
	"timestamp":"time.Time",
	"tinyblob":"[]byte",
	"tinyint":"int",
	"varbinary":"[]byte",
	"varchar":"string",
	"year":"time.Time",
}

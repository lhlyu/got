package g

type Table struct {
	Name   string
	Columns []*Column
}

type Column struct {
	Name   string
	Type   string
	ColumnName string
	JsonName string
	ColumnComment  string
}

type ColumnAttr struct {
	NameLen  int
	TypeLen  int
	ColumnLen int
	JsonLen  int
}
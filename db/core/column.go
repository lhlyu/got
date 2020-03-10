package core

type Column struct {
	Name              string
	TableName         string
	Sort              int
	Default           string
	IsNull            bool
	DataType          string
	MaxLength         int // 字符最大长度
	OctetLength       int // 数据的存储长度
	IsNumber          bool
	IsUnsigned        bool
	NumPrecision      int
	NumScale          int
	DatetimePrecision int
	ColumnType        string
	ColumnKey         string
	Extra             string
	Comment           string
}

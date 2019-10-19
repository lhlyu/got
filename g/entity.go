package g

import "github.com/lhlyu/got/qa"

type Column struct {
	Name      string
	Type      string
	TitleName string
	SmallName string
	Comment   string
}

type ColumnLen struct {
	NameLen      int
	TitleNameLen int
	SmallNameLen int
	TypeLen      int
}

type Tag struct {
	Name string
	T    int // 0 - 原生  1 - 大驼峰  2 - 小驼峰
	F    func(string) string
}

type Row struct {
	S  string
	C  string
	B  bool
}

type Tab struct {
	T   string
	B   bool
}

var tagMap = map[string]*Tag{
	qa.OptionJson:         &Tag{"json", 2, toSmallHump},
	qa.OptionDb:           &Tag{"db", 0, nil},
	qa.OptionXorm:         &Tag{"xorm", 0, nil},
	qa.OptionGorm:         &Tag{"gorm", 0, nil},
	qa.OptionValid:        &Tag{"valid", 2, toSmallHump},
	qa.OptionPlaceholder1: &Tag{"###", 2, toSmallHump},
	qa.OptionPlaceholder2: &Tag{"@@@", 2, toSmallHump},
}

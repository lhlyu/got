package gen

import (
	"bytes"
	"fmt"
	"github.com/lhlyu/got/db/core"
	"github.com/lhlyu/yutil/v2"
	"go/format"
)

type Gen struct {
	tagHandler  func(col *core.Column) string
	funcHandler func(tab *core.Table) string
}

func NewGen(tagHandler func(col *core.Column) string, funcHandler func(tab *core.Table) string) Gen {
	return Gen{
		tagHandler:  tagHandler,
		funcHandler: funcHandler,
	}
}

func (g Gen) ToStruct(dict map[string]string, tabs ...*core.Table) map[string]string {
	m := make(map[string]string)
	for _, tab := range tabs {
		buf := bytes.Buffer{}
		if tab.Comment != "" {
			buf.WriteString(fmt.Sprintf("// %s\n", tab.Comment))
		}
		buf.WriteString(fmt.Sprintf("type %s struct {\n", yutil.String.BigCamelCase(tab.Name)))
		for _, col := range tab.Columns {
			fieldName := yutil.String.BigCamelCase(col.Name)
			fieldType := dict[col.DataType]
			if col.IsUnsigned {
				if v, ok := dict[col.DataType+" unsigned"]; ok {
					fieldType = v
				}
			}
			// tinyint 存在无符
			fieldTag := ""
			if g.tagHandler != nil {
				fieldTag = g.tagHandler(col)
			}
			fieldComment := ""
			if col.Comment != "" {
				fieldComment = "// " + col.Comment
			}
			buf.WriteString(fmt.Sprintf("%s %s %s %s\n", fieldName, fieldType, fieldTag, fieldComment))
		}
		buf.WriteString("}\n")
		if g.funcHandler != nil {
			buf.WriteString(g.funcHandler(tab))
		}
		bts, _ := format.Source(buf.Bytes())
		fmt.Println(string(bts))
		m[tab.Name] = string(bts)
	}
	return m
}

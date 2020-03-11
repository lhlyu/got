package gen

import (
	"bytes"
	"fmt"
	"github.com/lhlyu/got/v2/db/core"
	"github.com/lhlyu/yutil/v2"
	"go/format"
	"strings"
)

type Gen struct {
	tagHandlers  []func(col *core.Column) string
	funcHandlers []func(tab *core.Table) string
}

func NewGen() *Gen {
	return &Gen{}
}

func (g *Gen) AddTagHandlers(tagHandlers ...func(col *core.Column) string) *Gen {
	g.tagHandlers = append(g.tagHandlers, tagHandlers...)
	return g
}

func (g *Gen) AddFuncHandlers(funcHandlers ...func(tab *core.Table) string) *Gen {
	g.funcHandlers = append(g.funcHandlers, funcHandlers...)
	return g
}

func (g *Gen) ToStruct(dict map[string]string, tabs ...*core.Table) map[string]string {
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
			if g.tagHandlers != nil {
				var tags []string
				for _, th := range g.tagHandlers {
					tags = append(tags, th(col))
				}
				fieldTag = fmt.Sprintf("`%s`", strings.Join(tags, " "))
			}
			fieldComment := ""
			if col.Comment != "" {
				fieldComment = "// " + col.Comment
			}
			buf.WriteString(fmt.Sprintf("%s %s %s %s\n", fieldName, fieldType, fieldTag, fieldComment))
		}
		buf.WriteString("}\n")
		if g.funcHandlers != nil {
			for _, fh := range g.funcHandlers {
				buf.WriteString(fh(tab))
			}

		}
		bts, _ := format.Source(buf.Bytes())
		fmt.Println(string(bts))
		m[tab.Name] = string(bts)
	}
	return m
}

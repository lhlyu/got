package g

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/lhlyu/got/db"
	"github.com/lhlyu/got/qa"
	"github.com/lhlyu/got/util"
	"regexp"
	"strconv"
	"strings"
)

const indent = "    "  // 缩进
const gap = 2

type gor struct {
	Tags      []string
	OutWay    string
	driverName string
	dir       string         // 输出文件夹地址
	name      func(columnName string) string  // 修改命名规则，默认是匈牙利命名(大驼峰)，下划线会自动去掉，例子如下
}
/**
默认命名
abcd    => Abcd
ab_cd   => AbCd
 */

func NewGor(tags []string,outway,driverName string,name func(columnName string) string) *gor{
	return &gor{tags,outway,driverName,"./",name}
}

func (g *gor) Run(ds []*db.TableInfo) {
	if len(ds) == 0{
		fmt.Println("没有找到表")
		return
	}
	typeMap := getTypeMap(g.driverName)

	tableMap := make(map[string][]*Column)
	tableLenMap := make(map[string]*ColumnLen)

	for _,d := range ds{
		cs,has := tableMap[d.TabelName]
		if !has{
			cs = make([]*Column,0)
		}
		c := &Column{
			Name:         d.ColumnName,
			Type:         getType(d.ColumnType,typeMap),
			TitleName:    toTitle(d.ColumnName),
			SmallName:    toSmallHump(d.ColumnName),
			Comment:      d.ColumnComment,
		}
		cs = append(cs,c)
		tableMap[d.TabelName] = cs

		le,has := tableLenMap[d.TabelName]
		if !has{
			tableLenMap[d.TabelName] = &ColumnLen{
				NameLen:      len(c.Name),
				TitleNameLen: len(c.TitleName),
				SmallNameLen: len(c.SmallName),
				TypeLen:      len(c.Type),
			}
		}else{
			if le.NameLen < len(c.Name){
				le.NameLen = len(c.Name)
			}
			if le.TitleNameLen < len(c.TitleName){
				le.TitleNameLen = len(c.TitleName)
			}
			if le.SmallNameLen < len(c.SmallName){
				le.SmallNameLen = len(c.SmallName)
			}
			if le.TypeLen < len(c.Type){
				le.TypeLen = len(c.Type)
			}
			tableLenMap[d.TabelName] = le
		}
	}

	// table handler

	tMap := make(map[string][]*Row)
	tlMap := make(map[string]int)
	for t,cs := range tableMap{

		le := tableLenMap[t]
		for _,c := range cs{
			tags := ""
			for _,tag := range g.Tags{
				if tf,has := tagMap[tag];has{
					switch tf.T {
					case 0:
						format := tf.Name + `:%-` + strconv.Itoa(le.NameLen + gap*2) + `s`
						tags += fmt.Sprintf(format,`"` + c.Name + `"`)
					case 1:
						format := tf.Name + `:%-` + strconv.Itoa(le.TitleNameLen + gap*2) + `s`
						tags += fmt.Sprintf(format,`"` + c.TitleName + `"`)
					case 2:
						format := tf.Name + `:%-` + strconv.Itoa(le.SmallNameLen + gap*2) + `s`
						tags += fmt.Sprintf(format,`"` + c.SmallName + `"`)
					}

				}
			}
			if len(tags) > 0{
				tags = strings.TrimRight(tags," ")
				tags = "`" + tags + "`"
			}
			format := "%s%-" + strconv.Itoa(le.TitleNameLen + gap) + "s%-" + strconv.Itoa(le.TypeLen) + "s  %s"
			buf := NewBufer()
			buf.Addf(format,indent,c.TitleName,c.Type,tags)


			tl,has := tlMap[t]
			if !has{
				tl = len(buf.String())
			}
			if tl < len(buf.String()){
				tl = len(buf.String())
			}
			tlMap[t] = tl


			ts,has := tMap[t]
			if !has{
				ts = make([]*Row,0)
			}
			row := &Row{
				S: buf.String(),
				C: c.Comment,
				B: c.Type == "time.Time",
			}
			ts = append(ts,row)
			tMap[t] = ts

		}

	}

	tabMap := make(map[string]*Tab)

	for k,value := range tMap{
		buf := NewBufer()
		buf.Addf("type %s struct {\n",toTitle(k))
		b := false
		for _,v := range value{
			format := "%-" +strconv.Itoa(tlMap[k] + gap) + "s"
			if v.C != ""{
				format += " // " + v.C
			}
			buf.Addf(format,v.S)
			buf.Add("\n")
			if v.B{
				b = true
			}
		}
		buf.Addf("}\n")
		tabMap[k] = &Tab{
			T: buf.String(),
			B: b,
		}
	}

	g.out(tabMap)
	return
}

func (g *gor) out(m map[string]*Tab){
	switch g.OutWay {
	case qa.RadioSingle:
		buf := NewBufer()
		b := false
		for _,v := range m{
			buf.Add(v.T)
			buf.Add("\n")
			if v.B{
				b = true
			}
		}
		content := "package model\n\n"
		if b{
			content += "import \"time\"\n\n"
		}
		content += buf.String()
		util.WriteFile(g.dir + "model.go",content)
		fmt.Printf("内容已写入到model.go文件\n")
	case qa.RadioMutil:
		for k,v := range m{
			buf := NewBufer()
			buf.Add("package model\n\n")
			if v.B{
				buf.Add("import \"time\"\n\n")
			}
			buf.Add(v.T)
			buf.Add("\n")
			util.WriteFile(g.dir + k + ".go",buf.String())
			fmt.Printf("表%s内容已写入到%s.go文件\n",k,k)
		}
	case qa.RadioClip:
		buf := NewBufer()
		for _,v := range m{
			buf.Add(v.T)
			buf.Add("\n")
		}
		clipboard.WriteAll(buf.String())
		fmt.Println("内容已写入剪贴板")
	}
	return
}

// 大驼峰
func toTitle(s string) string{
	s = strings.ReplaceAll(s,"_"," ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s," ","")
	return s
}

// 小驼峰
func toSmallHump(s string) string{
	s = strings.ReplaceAll(s,"_"," ")
	s = strings.Title(s)
	s = strings.ReplaceAll(s," ","")
	s = strings.ToLower(s[0:1]) + s[1:]
	return s
}

// 获取字段类型
func getType(s string,typeMap map[string]string) string{
	re, _ := regexp.Compile("\\(\\d+\\)")
	s = re.ReplaceAllString(s, "")
	if s,has := typeMap[s];has{
		return s
	}
	return "string"
}

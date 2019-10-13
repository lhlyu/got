package g

import (
	"github.com/lhlyu/got/db"
	"strings"
	"regexp"
	"strconv"
	"github.com/atotto/clipboard"
	"fmt"
	"github.com/lhlyu/got/util"
)

const indent = "    "  // 缩进
const gap = 2

type gor struct {
	hasJson   bool           // 是否生成json标签
 	hasDb     bool           // 是否生成db标签
	hasGSst   bool           // 是否生成get set方法
	isSingle  bool           // 所有struct写入到一个文件
	isMutil  bool            // 每个struct写入到一个文件
	isClip    bool
	driverName string
	dir       string         // 输出文件夹地址
	name      func(columnName string) string  // 修改命名规则，默认是匈牙利命名(大驼峰)，下划线会自动去掉，例子如下
}
/**
默认命名
abcd    => Abcd
ab_cd   => AbCd
 */

func NewGor(hasJson,hasDb,hasGSst,isSingle,isMutil,isClip bool,driverName,dir string,name func(columnName string) string) *gor{
	return &gor{hasJson,hasDb,hasGSst,isSingle,isMutil,isClip,driverName,dir,name}
}

func (g *gor) Run(ds []*db.TableInfo) {
	if len(ds) == 0{
		return
	}
	typeMap := getTypeMap(g.driverName)
	tabMap := make(map[string][]*Column)
	attrMap := make(map[string]*ColumnAttr)
	timeMap := make(map[string]bool)
	for _,d := range ds{
		t,has := tabMap[d.TabelName]
		if !has{
			t = make([]*Column,0)
		}
		c := &Column{
			ColumnName: d.ColumnName,
			ColumnComment: d.ColumnComment,
		}
		c.Type = getType(d.ColumnType,typeMap)
		c.JsonName = toSmallHump(d.ColumnName)
		if g.name != nil{
			c.Name = g.name(d.ColumnName)
		}else{
			c.Name = toTitle(d.ColumnName)
		}
		t = append(t,c)
		tabMap[d.TabelName] = t

		if c.Type == "time.Time"{
			timeMap[d.TabelName] = true
		}

		attr,ok := attrMap[d.TabelName]
		if !ok{
			attr = &ColumnAttr{}
		}
		if len(c.Name) > attr.NameLen{
			attr.NameLen = len(c.Name)
		}
		if len(c.Type) > attr.TypeLen{
			attr.TypeLen = len(c.Type)
		}
		if len(c.ColumnName) > attr.ColumnLen{
			attr.ColumnLen = len(c.ColumnName)
		}
		if len(c.JsonName) > attr.JsonLen{
			attr.JsonLen = len(c.JsonName)
		}
		attrMap[d.TabelName] = attr
	}
	modelMap := make(map[string]string)
	for key,value := range tabMap{
		attr := attrMap[key]
		format := "%s%-" + strconv.Itoa(attr.NameLen + gap) + "s%-" +  strconv.Itoa(attr.TypeLen + gap) +  "s"
		tag := ""
		if g.hasJson{
			tag += `json:%-` + strconv.Itoa(attr.JsonLen + gap) + `s`
		}
		if g.hasDb{
			if len(tag) > 0{
				tag += "  "
			}
			tag += `db:%-` + strconv.Itoa(attr.ColumnLen + gap) + `s`
		}
		if len(tag) > 0{
			format += "`" + tag + "`"
		}

		buf := NewBufer()
		buf.Addf("type %s struct {\n",toTitle(key))
		for _,v := range value{
			params := []interface{}{indent,v.Name,v.Type}
			if g.hasJson{
				params = append(params,`"` + v.JsonName+`"`)
			}
			if g.hasDb{
				params = append(params,`"` + v.ColumnName+`"`)
			}
			buf.Addf(format,params...)
			if len(v.ColumnComment) > 0{
				buf.Add("  // ",v.ColumnComment)
			}
			buf.Add("\n")
		}
		buf.Add("}\n")
		modelMap[key] = buf.String()
	}
	g.out(modelMap,timeMap)
	return
}

func (g *gor) out(m map[string]string,t map[string]bool){
	buf := NewBufer()
	for _,v := range m{
		buf.Add(v)
		buf.Add("\n")
	}
	if g.isClip{
		clipboard.WriteAll(buf.String())
		fmt.Println("内容已写入到剪贴板")
	}
	if g.isSingle{
		hasTimePackage := false
		for _,v := range t{
			if v{
				hasTimePackage = true
			}
		}
		content := "package model\n\n"
		if hasTimePackage{
			content += "import \"time\"\n\n"
		}
		content += buf.String()
		util.WriteFile(g.dir + "model.go",content)
		fmt.Printf("内容已写入到model.go文件\n")
	}
	if g.isMutil{
		for k,v := range m{
			hasTimePackage := t[k]
			content := "package model\n\n"
			if hasTimePackage{
				content += "import \"time\"\n\n"
			}
			content += v
			util.WriteFile(g.dir + k + ".go",content)
			fmt.Printf("表%s内容已写入到%s.go文件\n",k,k)
		}
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
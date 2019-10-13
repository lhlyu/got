package qa

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/lhlyu/got/util"
	"fmt"
)

// 用户交互
type Answer struct {
	Conf  string
	Host  string
	User  string
	Pass  string
	Port  string
	DB    string
	Table string
	Advance []string
	Save  bool
}

const Custom = "自定义"
const (
	OptionJson = "添加JSON标签"
	OptionDb = "添加DB标签"
	OptionSingle = "写入到单个文件"
	OptionMutil = "写入到多个文件"
	OptionClip = "写入到剪贴板"
)

func NewQuestions(flag int) []*survey.Question{
	if flag == 1{
		return defaultQ
	}
	if flag == 2{
		return configQ
	}
	fileList := util.GetAllFile("./","got-")
	if len(fileList) > 0{
		options := []string{Custom}
		options = append(options,fileList...)
		return UseConfig(options)
	}

	return defaultQ
}

func (self *Answer) ReadConf(){
	if self.Conf != ""{
		contents := util.ReadFileLinesTrim("./" + self.Conf)
		self.Host = contents[0]
		self.Port = contents[1]
		self.User = contents[2]
		util.AnyDecode(contents[3],&self.Pass)
		self.DB = contents[4]
	}
}

func (self *Answer) SaveConf(){
	fileName := fmt.Sprintf("./got-%s",self.DB)
	content := self.Host
	content += "\n" + self.Port
	content += "\n" + self.User
	content += "\n" + util.AnyEncode(self.Pass)
	content += "\n" + self.DB
	util.WriteFile(fileName,content)
}

var defaultQ = []*survey.Question{
	{
		Name:      "Host",
		Prompt:    &survey.Input{
			Message: "Host: ",
			Default:"localhost",
		},
		Validate:  survey.Required,
	},
	{
		Name:      "Port",
		Prompt:    &survey.Input{
			Message: "端口: ",
			Default:"3306",
		},
		Validate:  survey.Required,
	},
	{
		Name:      "User",
		Prompt:    &survey.Input{
			Message: "用户: ",
			Default:"root",
		},
		Validate:  survey.Required,
	},
	{
		Name:      "Pass",
		Prompt:    &survey.Input{
			Message: "密码: ",
			Default:"",
		},
		Validate:  survey.Required,
	},
	{
		Name:      "DB",
		Prompt:    &survey.Input{
			Message: "数据库: ",
			Default:"",
		},
		Validate:  survey.Required,
	},
	{
		Name:      "Table",
		Prompt:    &survey.Input{
			Message: "表(默认所有表): ",
			Default:"",
		},
	},
	{
		Name:      "Advance",
		Prompt: &survey.MultiSelect{
			Message: "高级设置(空格选择):",
			Options: []string{OptionJson, OptionDb, OptionMutil,OptionSingle,OptionClip},
			Default: []string{OptionClip},
		},
	},
	{
		Name:      "Save",
		Prompt:    &survey.Confirm{
			Message: "是否保存配置，方便下次连接?",
		},
		Validate:  survey.Required,
	},
}

var configQ = []*survey.Question{
	{
		Name:      "Table",
		Prompt:    &survey.Input{
			Message: "表(默认所有表): ",
			Default:"",
		},
	},
	{
		Name:      "Advance",
		Prompt: &survey.MultiSelect{
			Message: "高级设置(空格选择):",
			Options: []string{OptionJson, OptionDb, OptionMutil,OptionSingle,OptionClip},
			Default: []string{OptionClip},
		},
	},
}


func UseConfig(options []string) []*survey.Question{
	return []*survey.Question{
		{
			Name: "Conf",
			Prompt: &survey.Select{
				Message: "选择已有配置: ",
				Options: options,
				Default: options[0],
			},
			Validate: survey.Required,
		},
	}
}



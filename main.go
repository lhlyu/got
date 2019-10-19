package main

import (
	"github.com/AlecAivazis/survey/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lhlyu/got/db"
	"github.com/lhlyu/got/g"
	"github.com/lhlyu/got/qa"
	"log"
)

const version = "v1.1.0"

func main() {
	start()
}

func start(){
	q1 := qa.NewQuestions(0)
	answer := &qa.Answer{}
	err := survey.Ask(q1,answer)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if answer.Conf != "" && answer.Conf != qa.Custom{
		answer.ReadConf()
		q2 := qa.NewQuestions(2)
		err = survey.Ask(q2,answer)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}else if answer.Conf == qa.Custom{
		q3 := qa.NewQuestions(1)
		err = survey.Ask(q3,answer)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	if answer.Save{
		answer.SaveConf()
	}
	conf := db.NewDbConf(answer.Host,answer.Port,answer.User,answer.Pass,answer.DB,"",answer.Table)
	dao := db.NewDao("mysql")
	dao.SetDB(conf)
	tf := dao.QueryTables()
	gor := g.NewGor(answer.Tags,answer.OutWay,"mysql",nil)
	gor.Run(tf)
}

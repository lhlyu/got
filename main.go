package main

import (
	"github.com/AlecAivazis/survey/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lhlyu/got/db"
	"github.com/lhlyu/got/qa"
	"log"
	"github.com/lhlyu/got/g"
)

const version = "v1.0.1"

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
	hasJson := false
	hasDb := false
	isSingle := false
	isMutil := false
	isClip := false
	dir := ""
	for _,v := range answer.Advance{
		switch v {
		case qa.OptionJson:
			hasJson = true
		case qa.OptionDb:
			hasDb = true
		case qa.OptionSingle:
			dir = "./"
			isSingle = true
		case qa.OptionMutil:
			dir = "./"
			isMutil = true
		case qa.OptionClip:
			isClip = true
		}
	}
	gor := g.NewGor(hasJson,hasDb,false,isSingle,isMutil,isClip,"mysql",dir,nil)
	gor.Run(tf)
}

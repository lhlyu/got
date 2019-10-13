package qa

import (
	"bytes"
	"fmt"
	"github.com/lhlyu/got/util"
	"time"
)

const conf = "./got-%s"

func SaveConf(answer *Answer){
	buf := bytes.Buffer{}
	buf.WriteString(toConf(answer.Host))
	buf.WriteString(toConf(answer.User))
	buf.WriteString(toConf(util.AnyEncode(answer.Pass)))
	buf.WriteString(toConf(answer.Port))
	buf.WriteString(toConf(answer.DB))
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	filePath := fmt.Sprintf(conf,answer.DB)
	util.WriteFile(filePath,buf.String())
}

func toConf(s string) string{
	return fmt.Sprintf("%s\n",s)
}
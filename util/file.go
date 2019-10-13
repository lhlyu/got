package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"bufio"
	"io"
)

// 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 写入文件
func WriteFile(filePath string,content string){
	dir,_ := path.Split(filePath)
	if !PathExists(dir){
		os.MkdirAll(dir,os.ModePerm)
	}
	//将指定内容写入到文件中
	err := ioutil.WriteFile(filePath, []byte(content),os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}

// 获取所有文件
func GetAllFile(pathname string,filter string) []string {
	fileList := []string{}
	rd, _ := ioutil.ReadDir(pathname)
	if rd != nil{
		for _, fi := range rd {
			if !fi.IsDir() {
				if strings.Contains(fi.Name(),filter){
					fileList = append(fileList, fi.Name())
				}
			}
		}
	}
	return fileList
}

func ReadFileLinesTrim(filePath string) []string{
	var lines []string
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Println("filePath is err : ",err.Error())
		return lines
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		linebyt, _,err := br.ReadLine()
		if err == io.EOF{
			break
		}
		line := string(linebyt)
		line = strings.TrimSpace(line)
		if line == ""{
			continue
		}
		lines = append(lines,line)
	}
	return lines
}
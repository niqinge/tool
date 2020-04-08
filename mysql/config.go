package mysql

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type DbConfig struct {
	Url      string `json:"url" name:"地址"`
	User     string `json:"user" name:"用户名"`
	PassWord string `json:"pass_word" name:"密码"`
	DbName   string `json:"db_name" name:"数据库名"`
}


func NewDbConfig (dbName string) *DbConfig {
	result := &DbConfig{
		Url:      "127.0.0.1",
		User:     "root",
		PassWord: "123456",
		DbName:   dbName,
	}

	data, err := ioutil.ReadFile("/usr/local/.db/mysql.pas")
	if err != nil {
		fmt.Println("读取mysql密码文件出错:" + err.Error())
	} else {
		result.PassWord = strings.TrimSpace(string(data))
	}

	name, err := ioutil.ReadFile("/usr/local/.db/mysql.uname")
	if err != nil {
		fmt.Println("读取mysql用户名文件出错:" + err.Error())
	} else {
		result.User = strings.TrimSpace(string(name))
	}
	return result
}
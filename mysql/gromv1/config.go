package gromv1

import (
    "fmt"
    "io/ioutil"
    "strings"
    "time"
)

type DbConfig struct {
    Url             string        `json:"url" name:"地址"`
    User            string        `json:"user" name:"用户名"`
    Port            string        `json:"port" name:"端口"`
    PassWord        string        `json:"pass_word" name:"密码"`
    DbName          string        `json:"db_name" name:"数据库名"`
    MaxIdleConns    int           `json:"max_idle_conns"`    // MaxIdleConns sets the maximum number of connections in the idle connection pool.
    MaxOpenConns    int           `json:"max_open_conns"`    // MaxOpenConns sets the maximum number of open connections to the database.
    ConnMaxLifetime time.Duration `json:"conn_max_lifetime"` // ConnMaxLifetime sets the maximum amount of time a connection may be reused.
}

//
func NewLocalDBConf(dbName string) *DbConfig {
    result := &DbConfig{
        Url:      "127.0.0.1",
        User:     "root",
        PassWord: "123456",
        Port:     "3306",
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

// flag是否走默认dbname, true默认db, falsecong conf读dbname
func (conf *DbConfig) genUri(flag bool) string {
    if !flag {
        return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.PassWord, conf.Url, conf.Port, conf.DbName)
    }
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.PassWord, conf.Url, conf.Port, "information_schema")
}

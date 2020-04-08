package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func OpenGormConn(conf *DbConfig) *gorm.DB {
	uri := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.PassWord, conf.Url, conf.DbName)

	//db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/account?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour * 3)

	//defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}

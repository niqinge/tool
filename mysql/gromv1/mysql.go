package gromv1

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "time"
)

func OpenGormConn(conf *DbConfig) (*gorm.DB, error) {
    return openGormConn(conf, false)
}

func openGormConn(conf *DbConfig, flag bool) (*gorm.DB, error) {
    db, err := gorm.Open("mysql", conf.genUri(flag))
    if err != nil {
        return nil, fmt.Errorf("gorm open mysql , err:%s", err)
    }

    if conf.MaxIdleConns == 0 {
        db.DB().SetMaxIdleConns(10)
    } else {
        db.DB().SetMaxIdleConns(conf.MaxIdleConns)
    }
    if conf.MaxOpenConns == 0 {
        db.DB().SetMaxOpenConns(100)
    } else {
        db.DB().SetMaxOpenConns(conf.MaxOpenConns)
    }
    if conf.ConnMaxLifetime != 0 {
        db.DB().SetConnMaxLifetime(time.Hour * 3)
    } else {
        db.DB().SetConnMaxLifetime(conf.ConnMaxLifetime)
    }

    // defer db.Close()
    if err := db.DB().Ping(); err != nil {
        return nil, fmt.Errorf("gorm ping mysql , err:%s", err)
    }
    return db, nil
}

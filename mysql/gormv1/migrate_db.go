package gormv1

import (
    "errors"
    "fmt"
)

type MigrateOrm struct {
    conf   *DbConfig
    models []interface{}
}

func NewMigrateOrm(conf *DbConfig) *MigrateOrm {
    return &MigrateOrm{conf: conf}
}

func (o *MigrateOrm) SetModels(ms ...interface{}) {
    o.models = append(o.models, ms...)
}

func (o *MigrateOrm) CreateDB() error {
    createDbSQL := "CREATE DATABASE IF NOT EXISTS " + o.conf.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
    orm, err := openGormConn(o.conf, true)
    if err != nil {
        return err
    }
    defer orm.Close()

    if err = orm.Exec(createDbSQL).Error; err != nil {
        return fmt.Errorf("create db, conf:%+v, err:%s", o.conf, err)
    }
    return nil
}

func (o *MigrateOrm) dropDB() error {
    dropDbSQL := "DROP DATABASE IF EXISTS " + o.conf.DbName + ";"
    orm, err := OpenGormConn(o.conf)
    if err != nil {
        return err
    }
    defer orm.Close()

    if err = orm.Exec(dropDbSQL).Error; err != nil {
        return fmt.Errorf("drop db, conf:%+v, err:%s", o.conf, err)
    }
    return nil
}

func (o *MigrateOrm) MigrateDB() error {
    orm, err := OpenGormConn(o.conf)
    if err != nil {
        return err
    }
    defer orm.Close()

    if len(o.models) == 0 {
        return errors.New("migrate db is not found model. ")
    }

    if err = orm.AutoMigrate(o.models...).Error; err != nil {
        return fmt.Errorf("migrate db, err:%s", err)
    }

    return nil
}

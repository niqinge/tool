package gorm

import (
    "errors"
    "fmt"
    "gorm.io/gorm"
)

type PageParam struct {
    Size    uint32 // 每页大小
    PageNo  uint32 // 页码
    OrderBy string // 排序字段
    IsDesc  bool   // 是的倒叙
}

type Base struct {
    db            *gorm.DB
    pageSizeLimit uint32 // 每页数据最大数
}

func NewBase(db *gorm.DB, pageSizeLimit uint32) *Base {
    if pageSizeLimit == 0 {
        pageSizeLimit = 1000
    }

    return &Base{db: db, pageSizeLimit: pageSizeLimit}
}

type IBase interface {
    GetByID(id int64, dbEntity interface{}) error

    Delete(ids []int64, dbEntity interface{}) error

    // UpdateBy wr 过滤条件， newData新数据
    UpdateBy(tableName string, wr, newData map[string]interface{}) error

    Create(dbEntity interface{}) error

    // QueryCountBy 查询数据库 统计个数
    QueryCountBy(tableName string, wr map[string]interface{}) (total int64, err error)

    UpdatesByIds(tableName string, ids []int64, updateInfo map[string]interface{}) error

    // QueryBy 查询数据库
    QueryBy(tableName string, pageParam *PageParam, wr map[string]interface{}, result interface{}) (err error)
}

// 查询数据库 统计个数
func (b *Base) QueryCountBy(tableName string, wr map[string]interface{}) (total int64, err error) {
    tmpDb := b.db.Table(tableName)
    if len(wr) > 0 {
        tmpDb = tmpDb.Where(wr)
    }
    err = tmpDb.Count(&total).Error
    return
}

func (b *Base) GetByID(id int64, dbEntity interface{}) error {
    return b.db.Where("id = ?", id).Last(dbEntity).Error
}

func (b *Base) Delete(ids []int64, dbEntity interface{}) error {
    return b.db.Where("id in ?", ids).Delete(dbEntity).Error
}

func (b *Base) UpdateBy(tableName string, wr, newData map[string]interface{}) error {
    if len(wr) == 0 {
        return fmt.Errorf("暂不支持全局更新%s表", tableName)
    }

    if len(newData) == 0 {
        return errors.New("未填需要的更新字段信息")
    }

    return b.db.Table(tableName).Where(wr).Updates(newData).Error
}

func (b *Base) Create(dbEntity interface{}) error {
    return b.db.Create(dbEntity).Error
}

func (b *Base) UpdatesByIds(tableName string, ids []int64, updateInfo map[string]interface{}) error {

    return b.db.Table(tableName).Where("id IN (?)", ids).Updates(updateInfo).Error
}

func (b *Base) QueryBy(tableName string, pageParam *PageParam, wr map[string]interface{}, result interface{}) (err error) {
    if pageParam == nil {
        return errors.New("页码参数必填")
    }
    if pageParam.Size > b.pageSizeLimit {
        return errors.New("每页查询数据不能大于设置最大数量")
    }

    tx := b.db.Table(tableName)

    if len(wr) > 0 {
        tx = tx.Where(wr)
    }

    if len(pageParam.OrderBy) > 0 {
        desc := "DESC"
        if !pageParam.IsDesc {
            desc = "ASC"
        }
        tx = tx.Order(fmt.Sprintf("%s %s", pageParam.OrderBy, desc))
    }

    tx = tx.Limit(int(pageParam.Size)).Offset(int(pageParam.PageNo * pageParam.Size))

    return tx.Find(result).Error
}

package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type base struct {
	ID   uint   `gorm:"column:id;index;primary_key;" `
	Name string `gorm:"name"`
}

func TestOpenGormConn(t *testing.T) {

	conf := NewDbConfig("mysql")

	db := OpenGormConn(conf)
	b := &base{Name: "zhangsan"}
	assert.NoError(t, db.CreateTable(b).Error)
	assert.NoError(t, db.Create(b).Error)
	assert.NoError(t, db.DropTable(b).Error)
}

package mysql

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type base struct {
	ID   uint   `gorm:"column:id;index;primary_key;" `
	Name string `gorm:"name"`
}

func TestOpenGormConn(t *testing.T) {

	conf := NewLocalDBConf("mysql")

	db, err := OpenGormConn(conf)
	require.NoError(t, err)
	b := &base{Name: "zhangsan"}
	assert.NoError(t, db.CreateTable(b).Error)
	assert.NoError(t, db.Create(b).Error)
	assert.NoError(t, db.DropTable(b).Error)
}

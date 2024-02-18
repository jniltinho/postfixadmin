package model

import (
	"time"
)

// Alias Postfix Admin - Virtual Aliases
type Alias struct {
	Address  string    `gorm:"primaryKey;column:address" json:"-"`
	Goto     string    `gorm:"column:goto" json:"goto"`
	Domain   string    `gorm:"column:domain" json:"domain"`
	Created  time.Time `gorm:"column:created" json:"created"`
	Modified time.Time `gorm:"column:modified" json:"modified"`
	Active   bool      `gorm:"column:active" json:"active"`
}

// TableName get sql table name.获取数据库表名
func (m *Alias) TableName() string {
	return "alias"
}

package model

import (
	"time"
)

// Log Postfix Admin - Log
type Log struct {
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
	Username  string    `gorm:"column:username" json:"username"`
	Domain    string    `gorm:"column:domain" json:"domain"`
	Action    string    `gorm:"column:action" json:"action"`
	Data      string    `gorm:"column:data" json:"data"`
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *Log) TableName() string {
	return "log"
}

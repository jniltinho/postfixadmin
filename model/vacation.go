package model

import (
	"time"
)

// Vacation Postfix Admin - Virtual Vacation
type Vacation struct {
	Email        string    `gorm:"primaryKey;column:email" json:"-"`
	Subject      string    `gorm:"column:subject" json:"subject"`
	Body         string    `gorm:"column:body" json:"body"`
	Cache        string    `gorm:"column:cache" json:"cache"`
	Domain       string    `gorm:"column:domain" json:"domain"`
	Created      time.Time `gorm:"column:created" json:"created"`
	Active       bool      `gorm:"column:active" json:"active"`
	Modified     time.Time `gorm:"column:modified" json:"modified"`
	Activefrom   time.Time `gorm:"column:activefrom" json:"activefrom"`
	Activeuntil  time.Time `gorm:"column:activeuntil" json:"activeuntil"`
	IntervalTime int       `gorm:"column:interval_time" json:"intervalTime"`
}

// TableName get sql table name.获取数据库表名
func (m *Vacation) TableName() string {
	return "vacation"
}

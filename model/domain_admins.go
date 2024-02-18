package model

import (
	"time"
)

// DomainAdmins Postfix Admin - Domain Admins
type DomainAdmins struct {
	Username string    `gorm:"column:username" json:"username"`
	Domain   string    `gorm:"column:domain" json:"domain"`
	Created  time.Time `gorm:"column:created" json:"created"`
	Active   bool      `gorm:"column:active" json:"active"`
	ID       int       `gorm:"primaryKey;column:id" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *DomainAdmins) TableName() string {
	return "domain_admins"
}

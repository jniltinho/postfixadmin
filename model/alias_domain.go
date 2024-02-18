package model

import (
	"time"
)

// AliasDomain Postfix Admin - Domain Aliases
type AliasDomain struct {
	AliasDomain  string    `gorm:"primaryKey;column:alias_domain" json:"-"`
	TargetDomain string    `gorm:"column:target_domain" json:"targetDomain"`
	Created      time.Time `gorm:"column:created" json:"created"`
	Modified     time.Time `gorm:"column:modified" json:"modified"`
	Active       bool      `gorm:"column:active" json:"active"`
}

// TableName get sql table name.获取数据库表名
func (m *AliasDomain) TableName() string {
	return "alias_domain"
}

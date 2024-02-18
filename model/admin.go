package model

import (
	"time"
)

// Admin Postfix Admin - Virtual Admins
type Admin struct {
	Username      string    `gorm:"primaryKey;column:username" json:"-"`
	Password      string    `gorm:"column:password" json:"password"`
	Created       time.Time `gorm:"column:created" json:"created"`
	Modified      time.Time `gorm:"column:modified" json:"modified"`
	Active        bool      `gorm:"column:active" json:"active"`
	Superadmin    bool      `gorm:"column:superadmin" json:"superadmin"`
	Phone         string    `gorm:"column:phone" json:"phone"`
	EmailOther    string    `gorm:"column:email_other" json:"emailOther"`
	Token         string    `gorm:"column:token" json:"token"`
	TokenValidity time.Time `gorm:"column:token_validity" json:"tokenValidity"`
}

// TableName get sql table name.获取数据库表名
func (m *Admin) TableName() string {
	return "admin"
}

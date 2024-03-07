package model

import (
	"time"
)

// Domain Postfix Admin - Virtual Domains
type Domain struct {
	Domain         string    `gorm:"primaryKey;column:domain" json:"domain"`
	Description    string    `gorm:"column:description" json:"description"`
	Aliases        int       `gorm:"column:aliases" json:"aliases"`
	Mailboxes      int       `gorm:"column:mailboxes" json:"mailboxes"`
	Maxquota       int64     `gorm:"column:maxquota" json:"maxquota"`
	Quota          int64     `gorm:"column:quota" json:"quota"`
	Transport      string    `gorm:"column:transport" json:"transport"`
	Backupmx       bool      `gorm:"column:backupmx" json:"backupmx"`
	Created        time.Time `gorm:"column:created" json:"created"`
	Modified       time.Time `gorm:"column:modified" json:"modified"`
	Active         bool      `gorm:"column:active" json:"active"`
	PasswordExpiry int       `gorm:"column:password_expiry" json:"password_expiry"`
	DomainEncode   string    `gorm:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *Domain) TableName() string {
	return "domain"
}

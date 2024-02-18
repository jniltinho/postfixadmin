package model

import (
	"time"
)

// Mailbox Postfix Admin - Virtual Mailboxes
type Mailbox struct {
	Username       string    `gorm:"primaryKey;column:username" json:"-"`
	Password       string    `gorm:"column:password" json:"password"`
	Name           string    `gorm:"column:name" json:"name"`
	Maildir        string    `gorm:"column:maildir" json:"maildir"`
	Quota          int64     `gorm:"column:quota" json:"quota"`
	LocalPart      string    `gorm:"column:local_part" json:"localPart"`
	Domain         string    `gorm:"column:domain" json:"domain"`
	Created        time.Time `gorm:"column:created" json:"created"`
	Modified       time.Time `gorm:"column:modified" json:"modified"`
	Active         bool      `gorm:"column:active" json:"active"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	EmailOther     string    `gorm:"column:email_other" json:"emailOther"`
	Token          string    `gorm:"column:token" json:"token"`
	TokenValidity  time.Time `gorm:"column:token_validity" json:"tokenValidity"`
	PasswordExpiry time.Time `gorm:"column:password_expiry" json:"passwordExpiry"`
}

// TableName get sql table name.获取数据库表名
func (m *Mailbox) TableName() string {
	return "mailbox"
}

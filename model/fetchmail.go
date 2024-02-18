package model

import (
	"time"
)

// Fetchmail [...]
type Fetchmail struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`
	Mailbox        string    `gorm:"column:mailbox" json:"mailbox"`
	SrcServer      string    `gorm:"column:src_server" json:"srcServer"`
	SrcAuth        string    `gorm:"column:src_auth" json:"srcAuth"`
	SrcUser        string    `gorm:"column:src_user" json:"srcUser"`
	SrcPassword    string    `gorm:"column:src_password" json:"srcPassword"`
	SrcFolder      string    `gorm:"column:src_folder" json:"srcFolder"`
	PollTime       int       `gorm:"column:poll_time" json:"pollTime"`
	Fetchall       bool      `gorm:"column:fetchall" json:"fetchall"`
	Keep           bool      `gorm:"column:keep" json:"keep"`
	Protocol       string    `gorm:"column:protocol" json:"protocol"`
	Usessl         bool      `gorm:"column:usessl" json:"usessl"`
	ExtraOptions   string    `gorm:"column:extra_options" json:"extraOptions"`
	ReturnedText   string    `gorm:"column:returned_text" json:"returnedText"`
	Mda            string    `gorm:"column:mda" json:"mda"`
	Date           time.Time `gorm:"column:date" json:"date"`
	Sslcertck      bool      `gorm:"column:sslcertck" json:"sslcertck"`
	Sslcertpath    string    `gorm:"column:sslcertpath" json:"sslcertpath"`
	Sslfingerprint string    `gorm:"column:sslfingerprint" json:"sslfingerprint"`
	Domain         string    `gorm:"column:domain" json:"domain"`
	Active         bool      `gorm:"column:active" json:"active"`
	Created        time.Time `gorm:"column:created" json:"created"`
	Modified       time.Time `gorm:"column:modified" json:"modified"`
	SrcPort        int       `gorm:"column:src_port" json:"srcPort"`
}

// TableName get sql table name.获取数据库表名
func (m *Fetchmail) TableName() string {
	return "fetchmail"
}

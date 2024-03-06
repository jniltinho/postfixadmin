package model

import (
	"errors"
	"fmt"
	"postfixadmin/config"
	"postfixadmin/util"
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

func (m *Domain) CreateDomain() error {
	if config.DB().Model(m).Where("domain = ?", m.Domain).First(m).RowsAffected == 0 {
		return config.DB().Create(m).Error
	}
	message := fmt.Sprintf("Domain %s already exists", m.Domain)
	return errors.New(message)
}

func (m *Domain) ListDomains() []Domain {
	var domains []Domain
	config.DB().Where("domain != ?", "ALL").Order("domain").Find(&domains)

	for i, domain := range domains {
		domains[i].DomainEncode = util.URLEncode(domain.Domain)
	}
	return domains
}

func (m *Domain) GetDomain() error {
	err := config.DB().Where("domain = ?", m.Domain).First(m).Error
	m.DomainEncode = util.URLEncode(m.Domain)
	return err
}

func (m *Domain) DeleteDomain(domain string) error {
	return config.DB().Where("domain = ?", domain).Delete(m).Error
}

func (m *Domain) ActiveDomain(active int) int64 {
	// Active and Deactive domain
	// 1 = Active, 0 = Deactive
	return config.DB().Model(m).Update("active", active).RowsAffected
}

func (m *Domain) UpdateDomain() error {
	if !m.Active || !m.Backupmx {
		param := map[string]interface{}{"active": m.Active, "backupmx": m.Backupmx}
		config.DB().Model(m).Updates(param)
	}
	return config.DB().Updates(m).Error
}

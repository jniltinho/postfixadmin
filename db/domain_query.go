package db

import (
	"errors"
	"fmt"
	"postfixadmin/model"
	"postfixadmin/util"

	"gorm.io/gorm"
)

var FF = fmt.Sprintf

func CreateDomain(m *model.Domain) error {
	err := DB().First(m, "domain = ?", m.Domain).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return DB().Create(m).Error
	}
	message := FF("Domain %s already exists", m.Domain)
	return errors.New(message)
}

func ListDomains() []model.Domain {
	var m []model.Domain
	DB().Find(&m, "domain != ?", "ALL").Order("domain")

	for i, domain := range m {
		m[i].DomainEncode = util.URLEncode(domain.Domain)
	}
	return m
}

func GetDomain(domain string) (model.Domain, error) {
	var m model.Domain
	err := DB().First(&m, "domain = ?", domain).Error
	m.DomainEncode = util.URLEncode(m.Domain)
	return m, err
}

func DeleteDomain(domain string) error {
	var m model.Domain
	return DB().Delete(m, "domain = ?", domain).Error
}

func ActiveDomain(m *model.Domain, active int) int64 {
	// Active and Deactive domain
	// 1 = Active, 0 = Deactive
	return DB().Model(m).Update("active", active).RowsAffected
}

func UpdateDomain(m *model.Domain) error {
	if !m.Active || !m.Backupmx {
		param := map[string]interface{}{"active": m.Active, "backupmx": m.Backupmx}
		DB().Model(m).Updates(param)
	}
	return DB().Updates(m).Error
}

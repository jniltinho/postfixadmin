package model

import (
	"time"
)

// VacationNotification Postfix Admin - Virtual Vacation Notifications
type VacationNotification struct {
	OnVacation string    `gorm:"primaryKey;column:on_vacation" json:"-"`
	Vacation   Vacation  `gorm:"joinForeignKey:on_vacation;foreignKey:email;references:OnVacation" json:"vacationList"` // Postfix Admin - Virtual Vacation
	Notified   string    `gorm:"primaryKey;column:notified" json:"-"`
	NotifiedAt time.Time `gorm:"column:notified_at" json:"notifiedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *VacationNotification) TableName() string {
	return "vacation_notification"
}

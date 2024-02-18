package model

// Quota [...]
type Quota struct {
	Username string `gorm:"primaryKey;column:username" json:"-"`
	Path     string `gorm:"primaryKey;column:path" json:"-"`
	Current  int64  `gorm:"column:current" json:"current"`
}

// TableName get sql table name.获取数据库表名
func (m *Quota) TableName() string {
	return "quota"
}

package model

// Config PostfixAdmin settings
type Config struct {
	ID    int    `gorm:"primaryKey;column:id" json:"-"`
	Name  string `gorm:"column:name" json:"name"`
	Value string `gorm:"column:value" json:"value"`
}

// TableName get sql table name.获取数据库表名
func (m *Config) TableName() string {
	return "config"
}

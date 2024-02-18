package model

// Quota2 [...]
type Quota2 struct {
	Username string `gorm:"primaryKey;column:username" json:"-"`
	Bytes    int64  `gorm:"column:bytes" json:"bytes"`
	Messages int    `gorm:"column:messages" json:"messages"`
}

// TableName get sql table name.获取数据库表名
func (m *Quota2) TableName() string {
	return "quota2"
}

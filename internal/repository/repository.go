package repository

import (
	"fmt"
	"postfixadmin/log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) DB() *gorm.DB {
	return r.db
}

func NewDB(conf *viper.Viper) *gorm.DB {
	DSN := GetMysqlDSN(conf)
	maxIdleConns := conf.GetInt("mysql.max_idle_conns")
	maxOpenConns := conf.GetInt("mysql.max_open_conns")

	lv := logger.Error
	if conf.GetString("log.log_level") == "debug" {
		lv = logger.Info
	}
	cfg := &gorm.Config{Logger: logger.Default.LogMode(lv)}

	db, err := gorm.Open(mysql.Open(DSN), cfg)
	if err != nil {
		log.ERROR("Cannot connect to database %s", err.Error())
	}

	stdDB, _ := db.DB()
	stdDB.SetMaxIdleConns(maxIdleConns)
	stdDB.SetMaxOpenConns(maxOpenConns)
	stdDB.SetConnMaxLifetime(time.Hour)

	//db = db.Debug()
	return db
}

func GetMysqlDSN(conf *viper.Viper) string {

	dbUser := conf.GetString("mysql.user")
	dbPassword := conf.GetString("mysql.password")
	dbHost := conf.GetString("mysql.host")
	dbName := conf.GetString("mysql.dbname")

	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&timeout=10s&parseTime=True", dbUser, dbPassword, dbHost, dbName)

	return DSN
}

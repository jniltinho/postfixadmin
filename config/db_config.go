package config

import (
	"log/slog"
	"sync"
	"time"

	"postfixadmin/util"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() any
	instance    any
	once        sync.Once
	Conf        *viper.Viper
}

var dbInstance *DBInstance

// Instance gets the singleton instance
func (i *DBInstance) Instance() any {
	i.once.Do(func() { i.instance = i.initializer() })
	return i.instance
}

func (config *DBInstance) dbInit() any {
	conf := config.Conf
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
		//slog.Fatal("Cannot connect to database", zap.Error(err))
		slog.Error("Cannot connect to database", err)
	}

	stdDB, _ := db.DB()
	stdDB.SetMaxIdleConns(maxIdleConns)
	stdDB.SetMaxOpenConns(maxOpenConns)
	stdDB.SetConnMaxLifetime(time.Hour)

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func InitDBConnection(conf *viper.Viper) {
	config := &DBInstance{Conf: conf}
	dbInstance = &DBInstance{initializer: config.dbInit}
	util.LOG("Database Initializer")

	// Create Default Tables if not exists
	CreateSchema()
}

package database

import (
	"sync"

	"github.com/jniltinho/postfixadmin/config"
	"github.com/jniltinho/postfixadmin/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	initializer func() any
	instance    any
	once        sync.Once
}

var dbInstance *DBInstance

// Instance gets the singleton instance
func (i *DBInstance) Instance() any {
	i.once.Do(func() { i.instance = i.initializer() })
	return i.instance
}

func dbInit(conf *viper.Viper, zlog *log.Logger) any {

	DSN := config.GetMysqlDSN(conf)

	maxIdleConns := conf.GetInt("mysql.max_idle_conns")
	maxOpenConns := conf.GetInt("mysql.max_open_conns")

	lv := logger.Error
	if conf.GetString("log.log_level") == "debug" {
		lv = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(lv)}

	db, err := gorm.Open(mysql.Open(DSN), cfg)
	if err != nil {
		zlog.Fatal("Cannot connect to database", zap.Error(err))
	}

	stdDB, _ := db.DB()
	stdDB.SetMaxIdleConns(maxIdleConns)
	stdDB.SetMaxOpenConns(maxOpenConns)

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func ConnectDb(conf *viper.Viper, zlog *log.Logger) {
	dbInstance = &DBInstance{initializer: func() any { return dbInit(conf, zlog) }}
	zlog.Info("Database Initializer")

	// Create Default Tables if not exists
	CreateSchema(conf, zlog)
}

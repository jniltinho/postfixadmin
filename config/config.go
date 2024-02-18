package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetMysqlDSN(conf *viper.Viper) string {

	dbUser := conf.GetString("mysql.user")
	dbPassword := conf.GetString("mysql.password")
	dbHost := conf.GetString("mysql.host")
	dbName := conf.GetString("mysql.dbname")

	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&timeout=10s&parseTime=True", dbUser, dbPassword, dbHost, dbName)

	return DSN
}

package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "local.toml", "config path, eg: -conf local.toml")
		flag.Parse()
	}
	if envConf == "" {
		envConf = "local.toml"
	}
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)

}
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

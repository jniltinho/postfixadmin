package app

import (
	"postfixadmin/config"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Config *viper.Viper
}

func NewAppConfig(config *viper.Viper) *AppConfig {
	return &AppConfig{Config: config}
}

func (a *AppConfig) AppRun() {
	conf := a.Config

	secret := conf.GetString("security.jwt.key")
	config.InitDBConnection(conf)
	// Echo instance
	app := echo.New()

	//Routes
	Routes(app, secret)

	// Start server
	host := conf.GetString("http.host")
	//app.Logger.Fatal(app.Start(host))

	if conf.GetBool("http.ssl") {
		app.Logger.Fatal(app.StartTLS(host, conf.GetString("http.cert"), conf.GetString("http.key")))
	} else {
		app.Logger.Fatal(app.Start(host))
	}
}

package app

import (
	"postfixadmin/database"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AppRun(conf *viper.Viper) {

	secret := conf.GetString("security.jwt.key")
	database.ConnectDb(conf)
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

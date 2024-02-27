package app

import (
	"net/http"

	"postfixadmin/database"

	"postfixadmin/handlers"

	"postfixadmin/log"

	"github.com/spf13/viper"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var staticFS = echo.WrapHandler(http.FileServer(http.FS(FS)))

func AppRun(conf *viper.Viper) {
	// Logger
	logger := log.NewLog(conf)

	secret := conf.GetString("security.jwt.key")

	database.ConnectDb(conf, logger)
	// Echo instance
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//app.Use(handlers.WithAuth)
	app.Use(handlers.CheckSession)
	app.Use(middleware.Recover())
	app.GET("/static/*", staticFS)

	// Routes
	app.GET("/", handlers.Home)
	app.GET("/home", handlers.Home)
	app.GET("/home2", handlers.HomeOld)

	app.GET("/ListDomain", handlers.ListDomain)
	//app.GET("/login", handlers.LoginOld)
	//app.GET("/logout", handlers.LogoutUser)

	adm := app.Group("/adm")
	adm.GET("/login", handlers.Login)
	adm.POST("/login", handlers.LoginUser)
	adm.GET("/logout", handlers.LogoutUser)

	//GetRoute = handlers.GetURL(app)
	//loginUrl := GetRoute["handlers.Login"]

	// Start server
	host := conf.GetString("http.host")
	//app.Logger.Fatal(app.Start(host))

	if conf.GetBool("http.ssl") {
		app.Logger.Fatal(app.StartTLS(host, conf.GetString("http.cert"), conf.GetString("http.key")))
	} else {
		app.Logger.Fatal(app.Start(host))
	}
}

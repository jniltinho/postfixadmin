package app

import (
	"net/http"

	"github.com/jniltinho/postfixadmin/database"
	"github.com/jniltinho/postfixadmin/handlers"
	"github.com/jniltinho/postfixadmin/log"
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
	app.GET("/login", handlers.Login)
	app.POST("/login", handlers.LoginUser)
	app.GET("/logout", handlers.LogoutUser)

	teste := app.Group("/adm")
	teste.GET("/login", handlers.Login2)

	// Start server
	host := conf.GetString("http.host")
	//app.Logger.Fatal(app.Start(host))

	if conf.GetBool("http.ssl") {
		app.Logger.Fatal(app.StartTLS(host, conf.GetString("http.cert"), conf.GetString("http.key")))
	} else {
		app.Logger.Fatal(app.Start(host))
	}
}

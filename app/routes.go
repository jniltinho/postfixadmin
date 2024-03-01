package app

import (
	"postfixadmin/handlers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *AppConfig) Routes() {
	app := a.App
	secret := a.Secret
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//app.Use(handlers.WithAuth)
	app.Use(handlers.CheckSession)
	app.Use(middleware.Recover())
	staticFS := echo.MustSubFS(FS, "static")
	app.StaticFS("/static", staticFS)

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
}

package app

import (
	"postfixadmin/handlers"
	web "postfixadmin/public"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
)

func (r *AppConfig) runRoutes() {
	app := r.Echo
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(r.Secret))))

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//app.Use(handlers.WithAuth)
	app.Use(handlers.CheckSession)
	app.Use(middleware.Recover())

	// Static files
	web.StaticFiles(app)

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

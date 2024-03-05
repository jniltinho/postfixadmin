package app

import (
	"postfixadmin/handler"
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
	app.Use(handler.CheckSession)
	app.Use(middleware.Recover())

	// Static files
	web.StaticFiles(app)

	// Routes
	app.GET("/", handler.Home)
	app.GET("/home", handler.Home)

	app.GET("/ListDomain", handler.ListDomain)
	app.GET("/AddDomain", handler.FormNewDomain)
	app.POST("/AddDomain", handler.NewDomain)
	app.DELETE("/DelDomain/:domain", handler.DeleteDomain)
	//app.GET("/login", handlers.LoginOld)
	//app.GET("/logout", handlers.LogoutUser)

	adm := app.Group("/adm")
	adm.GET("/login", handler.Login)
	adm.POST("/login", handler.LoginUser)
	adm.GET("/logout", handler.LogoutUser)
}

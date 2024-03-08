package app

import (
	"postfixadmin/handler"
	"postfixadmin/handler/domainHandler"
	"postfixadmin/public"

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
	public.StaticFiles(app)

	// Routes
	app.GET("/", handler.Home)
	app.GET("/home", handler.Home)

	// Domain routes
	app.GET("/ListDomain", domainHandler.ListDomain)
	app.GET("/AddDomain", domainHandler.FormNewDomain)
	app.POST("/NewDomain", domainHandler.NewDomain)
	app.DELETE("/DelDomain/:domain", domainHandler.DelDomain)
	app.GET("/EditDomain/:domain", domainHandler.EditDomain)
	app.POST("/PostEditDomain", domainHandler.PostEditDomain)
	app.GET("/ActDomain/:domain/:active", domainHandler.ActDomain)

	adm := app.Group("/adm")
	adm.GET("/login", handler.Login)
	adm.POST("/login", handler.LoginUser)
	adm.GET("/logout", handler.LogoutUser)
}

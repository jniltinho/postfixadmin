package main

import (
	"embed"
	"net/http"

	"github.com/jniltinho/postfixadmin/config"
	"github.com/jniltinho/postfixadmin/handlers"
	"github.com/jniltinho/postfixadmin/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

//go:embed static/*
var content embed.FS

var staticFS = echo.WrapHandler(http.FileServer(http.FS(content)))

func main() {

	conf := config.NewConfig()
	logger := log.NewLog(conf)

	// Echo instance
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/static/*", staticFS)

	// Routes
	app.GET("/", handlers.Home)
	app.GET("/home", handlers.Home)
	app.GET("/login", handlers.Login)

	// Start server
	portBind := conf.GetString("http.port")
	logger.Info("Server Start", zap.String("host", "http://127.0.0.1:"+portBind))
	app.Logger.Fatal(app.Start(":" + portBind))
}

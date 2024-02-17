package main

import (
	"embed"
	"net/http"

	"github.com/jniltinho/postfixadmin/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var content embed.FS

var staticFS = echo.WrapHandler(http.FileServer(http.FS(content)))

func main() {
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
	app.Logger.Fatal(app.Start(":8008"))
}

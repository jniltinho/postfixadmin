package public

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed static/*
var FS embed.FS

func StaticFiles(app *echo.Echo) {
	staticFS := echo.MustSubFS(FS, "static")
	app.StaticFS("/static", staticFS)
}

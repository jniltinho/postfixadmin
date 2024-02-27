package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

var GetRoutes = map[string]string{
	"HomeUrl":   "/",
	"LoginUrl":  "/adm/login",
	"LogoutUrl": "/adm/logout",
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func hxRedirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", to)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}

	return c.Redirect(http.StatusSeeOther, to)
}

func LOG(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	r := len(msg) + 2

	color.Green(strings.Repeat("-", r))
	color.Red(msg)
	color.Green(strings.Repeat("-", r))
}

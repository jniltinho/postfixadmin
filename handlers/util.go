package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

var GetRoutes = map[string]string{
	"HomeUrl":   "/",
	"LoginUrl":  "/adm/login",
	"LogoutUrl": "/adm/logout",
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func LOG(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	slog.Info(msg)
}

func hxRedirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", to)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}

	return c.Redirect(http.StatusSeeOther, to)

}

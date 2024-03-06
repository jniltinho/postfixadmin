package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var FF = fmt.Sprintf

var GetRoutes = map[string]string{
	"HomeUrl":   "/",
	"LoginUrl":  "/adm/login",
	"LogoutUrl": "/adm/logout",
}

func Render(c echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response().Writer)
}

func hxRedirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", to)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}

	return c.Redirect(http.StatusSeeOther, to)
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render2(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

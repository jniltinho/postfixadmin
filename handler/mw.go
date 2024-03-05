package handler

import (
	"postfixadmin/log"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "/static") {
			return next(c)
		}

		log.INFO("Middleware WithAuth called %s", c.Request().URL.Path)
		return next(c)
	}
}

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if strings.Contains(c.Request().URL.Path, GetRoutes["LoginUrl"]) || strings.Contains(c.Request().URL.Path, "/static") {
			return next(c)
		}

		sess, _ := session.Get("session", c)
		if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
			//return c.Redirect(http.StatusSeeOther, GetRoutes["LoginUrl"])
			return hxRedirect(c, GetRoutes["LoginUrl"])
		}

		log.INFO("User: %s is Authenticated Token: %s", sess.Values["username"], sess.Values["session_token"])

		return next(c)
	}
}

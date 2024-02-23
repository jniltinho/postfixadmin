package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.Contains(c.Request().URL.Path, "/static") {
			return next(c)
		}

		fmt.Println("Middleware WithAuth called", c.Request().URL.Path)
		return next(c)
	}
}

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if strings.Contains(c.Request().URL.Path, "/login") || strings.Contains(c.Request().URL.Path, "/static") {
			return next(c)
		}

		sess, _ := session.Get("session", c)
		if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		return next(c)
	}
}

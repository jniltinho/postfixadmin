package handlers

import (
	"fmt"
	"strings"

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

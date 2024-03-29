package handler

import (
	"postfixadmin/view"

	"github.com/labstack/echo/v4"
)

// Home is the handler for the home page
func Home(c echo.Context) error {
	return Render(c, view.Home())
}

// Login is the handler for the login page
func Login(c echo.Context) error {
	LoginUrl := GetRoutes["LoginUrl"]
	return Render(c, view.Login(LoginUrl))
}

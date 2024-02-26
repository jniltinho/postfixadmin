package handlers

import (
	"postfixadmin/view"

	"github.com/labstack/echo/v4"
)

// Home is the handler for the home page
func Home(c echo.Context) error {
	return render(c, view.HomeNew())
}

// Home is the handler for the home page
func HomeOld(c echo.Context) error {
	return render(c, view.Home())
}

// Login is the handler for the login page
func Login(c echo.Context) error {
	LoginUrl := GetRoutes["LoginUrl"]
	return render(c, view.Login(LoginUrl))
}

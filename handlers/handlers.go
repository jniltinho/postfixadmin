package handlers

import (
	"net/http"
	"postfixadmin/model"
	"postfixadmin/view"

	"github.com/labstack/echo/v4"
)

// Home is the handler for the home page
func Home(c echo.Context) error {
	return Render(c, http.StatusOK, view.Home())
}

// Login is the handler for the login page
func Login(c echo.Context) error {
	LoginUrl := GetRoutes["LoginUrl"]
	return Render(c, http.StatusOK, view.Login(LoginUrl))
}

func ListDomain(c echo.Context) error {
	d := new(model.Domain)
	allDomains := d.ListDomains()
	return Render(c, http.StatusOK, view.ListDomain(allDomains))
}

func AddDomain(c echo.Context) error {
	return Render(c, http.StatusOK, view.AddDomain())
}

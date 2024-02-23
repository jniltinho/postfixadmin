package handlers

import (
	"net/http"

	"github.com/jniltinho/postfixadmin/view"

	"github.com/labstack/echo/v4"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func HelloName(c echo.Context) error {
	title := "Hello:" + c.Param("name")
	name := c.Param("name")
	//return templates.Index(title, name).Render(context.Background(), c.Response().Writer)
	return render(c, view.Index(title, name))
}

func Pong(c echo.Context) error {
	//return templates.Index("Pong", "FuncPong").Render(context.Background(), c.Response().Writer)
	return render(c, view.Index("Pong", "FuncPong"))
}

func Home(c echo.Context) error {
	return render(c, view.Home())
}

func Login(c echo.Context) error {
	return render(c, view.Login())
}

func Login2(c echo.Context) error {
	return render(c, view.Login2())
}

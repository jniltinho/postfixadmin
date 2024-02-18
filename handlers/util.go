package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func SchemaHandler(schema []byte) {
	query := string(schema)
	println(query)
	//database.CreateSchema(query)
}
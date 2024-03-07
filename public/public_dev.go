//go:build dev
// +build dev

package public

import (
	"os"
	"postfixadmin/log"

	"github.com/labstack/echo/v4"
)

func StaticFiles(app *echo.Echo) {

	if _, err := os.Stat("../public/static"); !os.IsNotExist(err) {
		log.INFO("DEV_MODE_USING_FOLDER: ../public/static")
		app.Static("/static", "../public/static")
	}

	if _, err := os.Stat("public/static"); !os.IsNotExist(err) {
		log.INFO("DEV_MODE_USING_FOLDER: public/static")
		app.Static("/static", "public/static")
	}

}

package handlers

import (
	"fmt"
	"net/http"
	"postfixadmin/database"
	"postfixadmin/model"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/tredoe/osutil/v2/userutil/crypt/md5_crypt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

var GetRoutes = map[string]string{
	"HomeUrl":   "/",
	"LoginUrl":  "/adm/login",
	"LogoutUrl": "/adm/logout",
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func LOG(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	slog.Info(msg)
}

func hxRedirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", to)
		c.Response().WriteHeader(http.StatusSeeOther)
		return nil
	}

	return c.Redirect(http.StatusSeeOther, to)

}

func GeneratePass(password string) string {
	c := md5_crypt.New()
	hash, err := c.Generate([]byte(password), []byte(getSalt(password)))
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func getSalt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	salt := string(bytes)[len(string(bytes))-11:]
	s := md5_crypt.GetSalt()
	salt = fmt.Sprintf("%s%s", s.MagicPrefix, salt)
	return salt
}

func ComparePass(hashedPassword, password string) bool {
	c := md5_crypt.New()
	err := c.Verify(hashedPassword, []byte(password))
	//fmt.Println(err)
	return err == nil
}

func CheckLogin(login, password string) bool {

	var user model.Admin
	result := database.DB().Where("active = ?", true).First(&user, "username = ?", login)

	if result.RowsAffected == 0 {
		//fmt.Println("User not found")
		return false
	}

	// Compare given user password with stored in found user.
	if !ComparePass(user.Password, password) {
		//fmt.Println("User not found")
		return false
	}

	//fmt.Println("User found")
	return true
}

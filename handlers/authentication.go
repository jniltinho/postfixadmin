package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var users = map[string]string{
	"admin":  "admin",
	"admin1": "admin1",
}

type LoginUserRequest struct {
	Username string `json:"username" xml:"username" form:"username" query:"username"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

func LoginUser(c echo.Context) error {
	userData := new(LoginUserRequest)

	if err := c.Bind(userData); err != nil {
		c.Logger().Error("Failed to bind the data", err)
		return err
	}

	expectedPassword, ok := users[userData.Username]

	if !ok || expectedPassword != userData.Password {
		return c.Redirect(http.StatusUnauthorized, "/login")
	}

	sessionToken := uuid.NewString()
	MaxAge := 3600 // 1 hour

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   MaxAge,
		HttpOnly: true,
		Secure:   true,
	}

	sess.Values["session_token"] = sessionToken
	sess.Values["username"] = userData.Username
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")

}

func LogoutUser(c echo.Context) error {
	sess, _ := session.Get("session", c)
	// Revoke users authentication

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values["username"] = nil
	sess.Values["authenticated"] = false
	sess.Values["session_token"] = nil

	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/login")
}

package handler

import (
	"postfixadmin/db"
	"postfixadmin/log"
	"postfixadmin/model"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

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

	//expectedPassword, ok := users[userData.Username]
	ok := checkLogin(userData.Username, userData.Password)

	if !ok {
		log.INFO("Failed to authenticate user: %s", userData.Username)
		//return c.Redirect(http.StatusUnauthorized, "/adm/login")
		return Redirect(c, GetRoutes["LoginUrl"])
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

	log.INFO("User: %s is Authenticated Token: %s", userData.Username, sessionToken)

	sess.Values["session_token"] = sessionToken
	sess.Values["username"] = userData.Username
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())
	//return c.Redirect(http.StatusSeeOther, "/")
	return Redirect(c, "/")
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
	return Redirect(c, GetRoutes["LoginUrl"])
}

func checkLogin(login, password string) bool {
	var user model.Admin
	res := db.DB().Select("password").First(&user, "username = ? AND active = 1", login)

	if res.RowsAffected == 0 {
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

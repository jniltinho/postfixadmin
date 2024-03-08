package domainHandler

//https://blog.benoitblanchon.fr/django-htmx-messages-framework/

import (
	"fmt"
	"strconv"
	"strings"

	"postfixadmin/db"
	"postfixadmin/view/domainView"

	"postfixadmin/handler"
	"postfixadmin/model"
	"postfixadmin/util"
	"time"

	"github.com/labstack/echo/v4"
)

var FF = fmt.Sprintf

type DomainRequest struct {
	Domain         string `json:"domain" form:"domain"`
	Description    string `json:"description" form:"description"`
	Aliases        int    `json:"aliases" form:"aliases"`
	Mailboxes      int    `json:"mailboxes" form:"mailboxes"`
	Active         bool   `json:"active" form:"active"`
	PasswordExpiry int    `json:"password_expiry" form:"password_expiry"`
	Backupmx       bool   `json:"backupmx" form:"backupmx"`
	DomainEncode   string `json:"domain_encode" form:"domain_encode"`
}

func ListDomain(c echo.Context) error {
	return handler.Render(c, domainView.ListDomain(db.ListDomains()))
}

func FormNewDomain(c echo.Context) error {
	return handler.Render(c, domainView.AddDomain())
}

func NewDomain(c echo.Context) error {

	dR := new(DomainRequest)
	domain := new(model.Domain)

	if err := c.Bind(dR); err != nil {
		message := domainView.Messages{Message: "Failed to bind the data", Alert: "error"}
		return handler.Render(c, domainView.FlashMessage(message))
	}

	domain.Domain = strings.ToLower(dR.Domain)
	domain.Description = dR.Description
	domain.Aliases = dR.Aliases
	domain.Mailboxes = dR.Mailboxes
	domain.Maxquota = 10
	domain.Quota = 2048
	domain.Transport = "virtual"
	domain.Backupmx = dR.Backupmx
	domain.PasswordExpiry = dR.PasswordExpiry
	domain.Created = time.Now()
	domain.Modified = time.Now()
	domain.Active = dR.Active

	res := db.CreateDomain(domain)
	if res != nil {
		message := domainView.Messages{Message: res.Error(), Alert: "error"}
		return handler.Render(c, domainView.FlashMessage(message))
	}

	m := FF("Domain Created: %s", domain.Domain)
	message := domainView.Messages{Message: m, Alert: "success", ResetForm: true}

	return handler.Render(c, domainView.FlashMessage(message))
}

func DelDomain(c echo.Context) error {
	domain := util.URLDecode(c.Param("domain"))

	res := db.DeleteDomain(domain)
	if res != nil {
		message := domainView.Messages{Message: res.Error(), Alert: "error"}
		list := db.ListDomains()
		return handler.Render(c, domainView.ListDomainTable(list, message))
	}

	m := FF("Domain Deleted: %s", domain)
	message := domainView.Messages{Message: m, Alert: "warning"}

	list := db.ListDomains()
	return handler.Render(c, domainView.ListDomainTable(list, message))
}

func EditDomain(c echo.Context) error {
	domain, _ := db.GetDomain(util.URLDecode(c.Param("domain")))
	return handler.Render(c, domainView.EditDomain(&domain))
}

func PostEditDomain(c echo.Context) error {

	dR := new(DomainRequest)
	domain := new(model.Domain)

	if err := c.Bind(dR); err != nil {
		//fmt.Println("Failed to bind the data", err.Error())
		res := FF("Failed Update: %s, Error: %s", domain.Domain, err.Error())
		message := domainView.Messages{Message: res, Alert: "error"}
		return handler.Render(c, domainView.FlashMessage(message))
	}

	if dR.DomainEncode != "" {
		dR.Domain = util.URLDecode(dR.DomainEncode)
	}

	domain.Domain = strings.ToLower(dR.Domain)
	domain.Description = dR.Description
	domain.Aliases = dR.Aliases
	domain.Mailboxes = dR.Mailboxes
	domain.Maxquota = 10
	domain.Quota = 2048
	domain.Transport = "virtual"
	domain.Backupmx = dR.Backupmx
	domain.PasswordExpiry = dR.PasswordExpiry
	domain.Modified = time.Now()
	domain.Active = dR.Active

	res := db.UpdateDomain(domain)
	if res != nil {
		message := domainView.Messages{Message: res.Error(), Alert: "error"}
		return handler.Render(c, domainView.FlashMessage(message))
	}

	m := FF("Domain Update: %s", domain.Domain)
	message := domainView.Messages{Message: m, Alert: "success", Redirect: "/ListDomain"}
	return handler.Render(c, domainView.FlashMessage(message))
}

func ActDomain(c echo.Context) error {
	domain := new(model.Domain)

	domain.Domain = util.URLDecode(c.Param("domain"))
	res := FF("Domain: %s is Activated", domain.Domain)

	if c.Param("active") == "1" || c.Param("active") == "0" {
		active, _ := strconv.Atoi(c.Param("active"))
		db.ActiveDomain(domain, active)
	}

	if c.Param("active") == "0" {
		res = FF("Domain: %s is Deactivated", domain.Domain)
	}

	message := domainView.Messages{Message: res, Alert: "warning"}
	list := db.ListDomains()
	return handler.Render(c, domainView.ListDomainTable(list, message))

}

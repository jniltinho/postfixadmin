package domain

//https://blog.benoitblanchon.fr/django-htmx-messages-framework/

import (
	"fmt"
	"strings"

	"postfixadmin/handler"
	"postfixadmin/model"
	"postfixadmin/util"
	"postfixadmin/view"
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
	d := new(model.Domain)
	js := `{"messages":[{"message":"It works!","tags":"success"}]}`

	c.Response().Header().Set("HX-Trigger", string(js))

	return handler.Render(c, view.ListDomain(d.ListDomains()))
}

func FormNewDomain(c echo.Context) error {
	return handler.Render(c, view.AddDomain())
}

func NewDomain(c echo.Context) error {

	dR := new(DomainRequest)
	domain := new(model.Domain)

	if err := c.Bind(dR); err != nil {
		message := view.Messages{Message: "Failed to bind the data", Alert: "error"}
		return handler.Render(c, view.FlashMessage(message))
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

	res := domain.CreateDomain()
	if res != nil {
		message := view.Messages{Message: res.Error(), Alert: "error"}
		return handler.Render(c, view.FlashMessage(message))
	}

	m := FF("Domain Created: %s", domain.Domain)
	message := view.Messages{Message: m, Alert: "success", ResetForm: true}

	return handler.Render(c, view.FlashMessage(message))
}

func DeleteDomain(c echo.Context) error {
	domain := new(model.Domain)
	domain.Domain = c.Param("domain")

	res := domain.DeleteDomain(domain.Domain)
	if res != nil {
		message := view.Messages{Message: res.Error(), Alert: "error"}
		list := domain.ListDomains()
		return handler.Render(c, view.ListDomainTable(list, message))
	}

	m := FF("Domain Deleted: %s", domain.Domain)
	message := view.Messages{Message: m, Alert: "warning"}

	list := domain.ListDomains()
	return handler.Render(c, view.ListDomainTable(list, message))
}

func EditDomain(c echo.Context) error {
	d := new(model.Domain)
	domainDecode := util.URLDecode(c.Param("domain"))
	d.Domain = domainDecode
	d.GetDomain()
	return handler.Render(c, view.EditDomain(d))
}

func UpdateDomain(c echo.Context) error {

	dR := new(DomainRequest)
	domain := new(model.Domain)

	if err := c.Bind(dR); err != nil {
		//fmt.Println("Failed to bind the data", err.Error())
		res := FF("Failed Update: %s, Error: %s", domain.Domain, err.Error())
		message := view.Messages{Message: res, Alert: "error"}
		return handler.Render(c, view.FlashMessage(message))
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

	res := domain.UpdateDomain()
	if res != nil {
		//message := view.Messages{Message: res.Error(), Alert: "error"}
		return handler.Redirect(c, "/ListDomain")
	}

	m := FF("Domain Update: %s", domain.Domain)
	message := view.Messages{Message: m, Alert: "success", Redirect: "/ListDomain"}

	return handler.Render(c, view.FlashMessage(message))
}

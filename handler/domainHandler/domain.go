package domainHandler

//https://blog.benoitblanchon.fr/django-htmx-messages-framework/

import (
	"strconv"

	"postfixadmin/db"
	"postfixadmin/view/domainView"
	"postfixadmin/view/ui"

	"postfixadmin/handler"
	"postfixadmin/model"
	"postfixadmin/util"

	"github.com/labstack/echo/v4"
)

func ListDomain(c echo.Context) error {
	return handler.Render(c, domainView.ListDomain(db.ListDomains()))
}

func FormNewDomain(c echo.Context) error {
	return handler.Render(c, domainView.AddDomain())
}

func NewDomain(c echo.Context) error {
	request := new(DomainRequest)

	if err := c.Bind(request); err != nil {
		m := ui.Messages{Message: "Failed to bind the data", Alert: "error"}
		return handler.Render(c, ui.FlashMessage(m))
	}

	util := NewUtil(c, ui.Messages{})
	domain, res := util.domainSave(request, true)
	if res != nil {
		m := ui.Messages{Message: res.Error(), Alert: "error"}
		return handler.Render(c, ui.FlashMessage(m))
	}

	mf := FF("Domain Created: %s", domain.Domain)
	m := ui.Messages{Message: mf, Alert: "success", ResetForm: true}
	return handler.Render(c, ui.FlashMessage(m))
}

func DelDomain(c echo.Context) error {
	domain := util.URLDecode(c.Param("domain"))

	res := db.DeleteDomain(domain)
	if res != nil {
		util := NewUtil(c, ui.Messages{Message: res.Error(), Alert: "error"})
		return util.domainList()
	}

	m := FF("Domain Deleted: %s", domain)
	util := NewUtil(c, ui.Messages{Message: m, Alert: "warning"})
	return util.domainList()
}

func EditDomain(c echo.Context) error {
	domain, _ := db.GetDomain(util.URLDecode(c.Param("domain")))
	return handler.Render(c, domainView.EditDomain(&domain))
}

func UpdateDomain(c echo.Context) error {
	request := new(DomainRequest)

	if err := c.Bind(request); err != nil {
		mf := FF("Failed to bind the data: %s", err.Error())
		m := ui.Messages{Message: mf, Alert: "error"}
		return handler.Render(c, ui.FlashMessage(m))
	}

	util := NewUtil(c, ui.Messages{})
	domain, res := util.domainSave(request, false)
	if res != nil {
		m := ui.Messages{Message: res.Error(), Alert: "error"}
		return handler.Render(c, ui.FlashMessage(m))
	}

	mf := FF("Domain Update: %s", domain.Domain)
	m := ui.Messages{Message: mf, Alert: "success", Redirect: "/ListDomain"}
	return handler.Render(c, ui.FlashMessage(m))
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

	util := NewUtil(c, ui.Messages{Message: res, Alert: "warning"})
	return util.domainList()
}

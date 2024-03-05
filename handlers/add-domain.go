package handlers

//https://blog.benoitblanchon.fr/django-htmx-messages-framework/

import (
	"net/http"
	"postfixadmin/log"
	"postfixadmin/model"
	"postfixadmin/view"
	"time"

	"github.com/labstack/echo/v4"
)

type DomainRequest struct {
	Domain         string `json:"domain" form:"domain"`
	Description    string `json:"description" form:"description"`
	Aliases        int    `json:"aliases" form:"aliases"`
	Mailboxes      int    `json:"mailboxes" form:"mailboxes"`
	Active         bool   `json:"active" form:"active"`
	PasswordExpiry int    `json:"password_expiry" form:"password_expiry"`
	Backupmx       bool   `json:"backupmx" form:"backupmx"`
}

// ACreatedomain is the handler for the add domain page
func CreateDomain(c echo.Context) error {

	dReq := new(DomainRequest)
	d := new(model.Domain)

	if err := c.Bind(dReq); err != nil {
		//return c.String(http.StatusBadRequest, "Failed to bind the data")
		viewData := view.Messages{
			Message: "Failed to bind the data",
			Alert:   "error",
		}
		return Render(c, http.StatusOK, view.DomainForm(viewData))
	}

	//INSERT INTO `domain` (domain,description,aliases,mailboxes,maxquota,quota,transport,backupmx,active,created,modified,password_expiry)
	//VALUES ('globo.com','Globo','40','40','10','2048','virtual','0','1',now(),now(),'365')

	d.Domain = dReq.Domain
	d.Description = dReq.Description
	d.Aliases = dReq.Aliases
	d.Mailboxes = dReq.Mailboxes
	d.Maxquota = 10
	d.Quota = 2048
	d.Transport = "virtual"
	d.Backupmx = dReq.Backupmx
	d.PasswordExpiry = dReq.PasswordExpiry
	d.Created = time.Now()
	d.Modified = time.Now()
	d.Active = dReq.Active

	domain := d.Domain
	description := d.Description
	aliases := d.Aliases
	mailboxes := d.Mailboxes
	active := d.Active
	alias := d.Aliases
	expiry := d.PasswordExpiry

	log.DEBUG("Domain: %s, Description: %s, Aliases: %d, Mailboxes: %d, Active: %t, Aliases: %d, Password Expiry: %d",
		domain, description, aliases, mailboxes, active, alias, expiry,
	)

	res := d.CreateDomain()
	if res != nil {
		viewData := view.Messages{
			Message: res.Error(),
			Alert:   "error",
		}
		return Render(c, http.StatusOK, view.DomainForm(viewData))
	}

	viewData := view.Messages{
		Message: "Domain Created Successfully: " + d.Domain,
		Alert:   "success",
	}

	return Render(c, http.StatusOK, view.DomainForm(viewData))

	//return c.String(http.StatusOK, "Domain Created Successfully: "+domain)
}

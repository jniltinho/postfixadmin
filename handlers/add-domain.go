package handlers

//https://blog.benoitblanchon.fr/django-htmx-messages-framework/

import (
	"net/http"
	"postfixadmin/log"
	"postfixadmin/view"

	"github.com/labstack/echo/v4"
)

type DomainRequest struct {
	Domain         string `json:"domain" form:"domain"`
	Description    string `json:"description" form:"description"`
	Aliases        int    `json:"aliases" form:"aliases"`
	Mailboxes      int    `json:"mailboxes" form:"mailboxes"`
	Active         bool   `json:"active" form:"active"`
	DefaultAliases string `json:"default_aliases" form:"default_aliases"`
	PasswordExpiry string `json:"password_expiry" form:"password_expiry"`
}

// ACreatedomain is the handler for the add domain page
func CreateDomain(c echo.Context) error {

	domainData := new(DomainRequest)

	if err := c.Bind(domainData); err != nil {
		//return c.String(http.StatusBadRequest, "Failed to bind the data")
		viewData := view.Messages{
			Message: "Failed to bind the data",
			Alert:   "error",
		}
		return Render(c, http.StatusOK, view.DomainForm(viewData))
	}

	domain := domainData.Domain
	description := domainData.Description
	aliases := domainData.Aliases
	mailboxes := domainData.Mailboxes
	active := domainData.Active
	defaultAliases := domainData.DefaultAliases
	passwordExpiry := domainData.PasswordExpiry

	log.DEBUG("Domain: %s, Description: %s, Aliases: %d, Mailboxes: %d, Active: %t, Default Aliases: %s, Password Expiry: %s",
		domain, description, aliases, mailboxes, active, defaultAliases, passwordExpiry,
	)

	viewData := view.Messages{
		Message: "Domain Created Successfully: " + domain,
		Alert:   "success",
	}

	return Render(c, http.StatusOK, view.DomainForm(viewData))

	//return c.String(http.StatusOK, "Domain Created Successfully: "+domain)
}

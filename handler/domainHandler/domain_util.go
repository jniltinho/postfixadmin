package domainHandler

import (
	"fmt"
	"postfixadmin/db"
	"postfixadmin/handler"
	"postfixadmin/model"
	"postfixadmin/util"
	"postfixadmin/view/domainView"
	"postfixadmin/view/ui"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var FF = fmt.Sprintf

type Util struct {
	Ctx     echo.Context
	Message ui.Messages
}

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

func NewUtil(c echo.Context, m ui.Messages) *Util {
	return &Util{
		Ctx:     c,
		Message: m,
	}
}

func (u *Util) domainList() error {
	message := u.Message
	list := db.ListDomains()
	return handler.Render(u.Ctx, domainView.ListDomainTable(list, message))
}

func (u *Util) domainSave(r *DomainRequest, create bool) (*model.Domain, error) {
	domain := new(model.Domain)

	if r.DomainEncode != "" {
		r.Domain = util.URLDecode(r.DomainEncode)
	}

	domain.Domain = strings.ToLower(r.Domain)
	domain.Description = r.Description
	domain.Aliases = r.Aliases
	domain.Mailboxes = r.Mailboxes
	domain.Maxquota = 10
	domain.Quota = 2048
	domain.Transport = "virtual"
	domain.Backupmx = r.Backupmx
	domain.PasswordExpiry = r.PasswordExpiry
	domain.Created = time.Now()
	domain.Modified = time.Now()
	domain.Active = r.Active

	if create {
		return domain, db.DomainNew(domain)
	}

	return domain, db.DomainUpdate(domain)
}

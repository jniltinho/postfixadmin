package types

import (
	"time"

	"github.com/uptrace/bun"
)

// Domain Postfix Admin - Virtual Domains
type Domain struct {
	bun.BaseModel  `bun:"table:domain,alias:d"`
	Domain         string    `bun:"domain,pk" json:"domain"`
	Description    string    `bun:"description" json:"description"`
	Aliases        int       `bun:"aliases" json:"aliases"`
	Mailboxes      int       `bun:"mailboxes" json:"mailboxes"`
	Maxquota       int64     `bun:"maxquota" json:"maxquota"`
	Quota          int64     `bun:"quota" json:"quota"`
	Transport      string    `bun:"transport" json:"transport"`
	Backupmx       bool      `bun:"backupmx" json:"backupmx"`
	Active         bool      `bun:"active" json:"active"`
	PasswordExpiry int       `bun:"password_expiry" json:"password_expiry"`
	Created        time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created"`
	Modified       time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"modified"`
	DomainEncode   string    `bun:"-"`
}

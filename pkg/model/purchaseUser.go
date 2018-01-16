package model

import (
	"time"
)

type PurchaseUser struct {
	Id                  int       `xorm:"not null pk autoincr INT(11)"`
	Username            string    `xorm:"not null VARCHAR(180)"`
	UsernameCanonical   string    `xorm:"not null unique VARCHAR(180)"`
	Email               string    `xorm:"not null VARCHAR(180)"`
	EmailCanonical      string    `xorm:"not null unique VARCHAR(180)"`
	Enabled             int       `xorm:"not null TINYINT(1)"`
	Salt                string    `xorm:"VARCHAR(255)"`
	Password            string    `xorm:"not null VARCHAR(255)"`
	LastLogin           time.Time `xorm:"DATETIME"`
	ConfirmationToken   string    `xorm:"unique VARCHAR(180)"`
	PasswordRequestedAt time.Time `xorm:"DATETIME"`
	Roles               string    `xorm:"not null LONGTEXT"`
	FirstName           string    `xorm:"VARCHAR(255)"`
	LastName            string    `xorm:"VARCHAR(255)"`
	Language            string    `xorm:"VARCHAR(16)"`
	Photo               string    `xorm:"VARCHAR(255)"`
	Configs             string    `xorm:"LONGTEXT"`
	Mobile              string    `xorm:"unique VARCHAR(255)"`
	Birthday            time.Time `xorm:"DATE"`
	CreatedAt           time.Time `xorm:"not null DATETIME"`
	UpdatedAt           time.Time `xorm:"not null DATETIME"`
}

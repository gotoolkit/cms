package model

import (
	"time"
)

type PurchaseBizCategory struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	RootId    int       `xorm:"index INT(11)"`
	ParentId  int       `xorm:"index INT(11)"`
	Lft       int       `xorm:"not null INT(11)"`
	Rgt       int       `xorm:"not null INT(11)"`
	Lvl       int       `xorm:"not null INT(11)"`
	Slug      string    `xorm:"VARCHAR(255)"`
	Status    int       `xorm:"not null default 0 INT(11)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	CreatedBy string    `xorm:"not null VARCHAR(255)"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedBy string    `xorm:"not null VARCHAR(255)"`
}

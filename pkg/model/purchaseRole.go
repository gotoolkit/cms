package model

type PurchaseRole struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Name      string `xorm:"VARCHAR(255)"`
	Roles     string `xorm:"LONGTEXT"`
	Userid    int    `xorm:"not null index INT(11)"`
	Companyid int    `xorm:"not null index INT(11)"`
}

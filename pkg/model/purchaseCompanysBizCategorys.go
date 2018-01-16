package model

type PurchaseCompanysBizCategorys struct {
	Companyid  int `xorm:"not null pk index INT(11)"`
	Categoryid int `xorm:"not null pk index INT(11)"`
}

package model

import (
	"time"
)

type PurchaseCompany struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	NameCn      string    `xorm:"VARCHAR(255)"`
	NamePt      string    `xorm:"VARCHAR(255)"`
	Tel         string    `xorm:"VARCHAR(255)"`
	Address     string    `xorm:"VARCHAR(255)"`
	Website     string    `xorm:"VARCHAR(255)"`
	Logo        string    `xorm:"VARCHAR(255)"`
	BrNumber    string    `xorm:"VARCHAR(255)"`
	BrFile      string    `xorm:"VARCHAR(255)"`
	ApprovelBy  string    `xorm:"VARCHAR(255)"`
	ApprovelAt  time.Time `xorm:"DATETIME"`
	Employee    int       `xorm:"INT(11)"`
	MLicense    string    `xorm:"VARCHAR(255)"`
	AvgAge      string    `xorm:"VARCHAR(255)"`
	OpenTime    string    `xorm:"VARCHAR(255)"`
	Fax         string    `xorm:"VARCHAR(255)"`
	Email       string    `xorm:"VARCHAR(255)"`
	Photos      string    `xorm:"LONGTEXT"`
	ProductImgs string    `xorm:"LONGTEXT"`
}

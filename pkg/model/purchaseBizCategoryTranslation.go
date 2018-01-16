package model

type PurchaseBizCategoryTranslation struct {
	Id             int    `xorm:"not null pk autoincr INT(11)"`
	Locale         string `xorm:"not null unique(purchase_biz_category_translation_unique_translation) VARCHAR(255)"`
	Title          string `xorm:"VARCHAR(255)"`
	Description    string `xorm:"LONGTEXT"`
	TranslatableId int    `xorm:"unique(purchase_biz_category_translation_unique_translation) index INT(11)"`
}

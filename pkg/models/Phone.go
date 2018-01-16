package models

import (
	"time"
)

// Phone ...
type Phone struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	UserID    uint
	PhoneNum  string    `json:"phone_num"`
	Code      string    `json:"code"`
	SentAt    time.Time `json:"sent_at"`
}

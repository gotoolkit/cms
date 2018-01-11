package models

import "time"

// User ...
type User struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Username  string     `gorm:"size:255"`
	Password  string     `json:"password"`
	Age       int
	PhoneNum  string `json:"phone_num"`
}

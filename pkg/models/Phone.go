package models

import (
	"time"
)

// Phone ...
type Phone struct {
	UserID   uint
	PhoneNum string    `json:"phone_num"`
	Code     string    `json:"code"`
	SentAt   time.Time `json:"sent_at"`
}

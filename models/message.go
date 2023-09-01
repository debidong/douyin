package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MessageID  int64 `gorm:"unique;not null;primary_Key"`
	ToUserID   int64 `gorm:"not null"`
	FromUserID int64 `gorm:"not null"`
	Content    string
}

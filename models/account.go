package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string  `gorm:"unique;not null;primary_Key"`
	UserId      int64   `gorm:"unique;not null"`
	Password    string  `gorm:"not null"`
	Subscribers []*User `gorm:"many2many:user_subscribers"`
	Followers   []*User `gorm:"many2many:user_followers"`
}

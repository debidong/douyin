package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoID       int64  `gorm:"unique;not null;primary_key"`
	AuthorID      int64  `gorm:"not null"` // 关联作者的 ID
	PlayURL       string `gorm:"not null"`
	CoverURL      string `gorm:"not null"`
	FavoriteCount int64  `gorm:"default:0"`
	CommentCount  int64  `gorm:"default:0"`
	IsFavorite    bool   `gorm:"default:false"`
	Title         string `gorm:"not null"`
}

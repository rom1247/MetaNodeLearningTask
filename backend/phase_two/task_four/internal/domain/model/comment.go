package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId  uint `persistence:"index"`
	PostID  uint `persistence:"index"`
	Content string
}

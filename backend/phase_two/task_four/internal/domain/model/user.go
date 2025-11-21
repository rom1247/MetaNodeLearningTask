package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `persistence:"unique"`
	Password string    `persistence:"not null" json:"-"`
	Email    string    `persistence:"unique" validate:"email"`
	Posts    []Post    `json:"-"` //忽略输出
	Comments []Comment `json:"-"`
}

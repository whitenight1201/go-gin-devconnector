package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"name" binding:"required,min=3,max=32"`
	Useremail      string `json:"email" binding:"required,min=10,max=32"`
	Password       string `gorm:"-" binding:"required,min=6,max=32"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
}

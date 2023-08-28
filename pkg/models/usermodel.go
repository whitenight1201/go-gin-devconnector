package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"name" gorm:"size:255;not null; uique" binding:"required,min=3,max=32"`
	Useremail      string `json:"email" gorm:"size:255;not null; uique" binding:"required,min=10,max=32"`
	Password       string `gorm:"-" binding:"required,min=6,max=32"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
}

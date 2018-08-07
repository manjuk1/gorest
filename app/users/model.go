package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	PasswordHash string  `gorm:"column:password_digest;not null"`
}

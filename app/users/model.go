package users

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"column:username"`
	Email        string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password_digest;not null"`
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *User) checkPassword(pwd string) error {
	password := []byte(pwd)
	hashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

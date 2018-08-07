package users

import (
	// "github.com/gosimple/slug"
	// "github.com/wangzitian0/golang-gin-starter-kit/common"
	// "github.com/wangzitian0/golang-gin-starter-kit/users"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/common"
)

type UserModelValidator struct {
	User struct {
		UserName     string  `json:"name" binding:"exists,required"`
		Email        string  `json:"email" binding:"exists,required"`
		PasswordHash string  `json:"password" binding:"exists,required"`
	} `json:"user"`
	userModel User `json:"-"`
}

func NewUserModelValidator() UserModelValidator {
	return UserModelValidator{}
}

func (s *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.userModel.UserName = s.User.UserName
	s.userModel.Email = s.User.Email
	s.userModel.SetPassword(s.User.PasswordHash)
	return nil
}

type LoginValidator struct {
	User struct {
		Email        string  `json:"email" binding:"exists,required"`
		Password 	 string  `json:"password" binding:"exists,required"`
	} `json:"user"`
	userDetails User `json:"-"`
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}

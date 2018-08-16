package users

import (
	"github.com/manjuk1/gorest/common"
)

type UserSerializer struct {
	user User
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:token`
}

func (self *UserSerializer) Response() UserResponse {
	user := UserResponse{
		ID:       self.user.Model.ID,
		Username: self.user.UserName,
		Email:    self.user.Email,
		Token:    common.GenToken(self.user.Model.ID),
	}
	return user
}

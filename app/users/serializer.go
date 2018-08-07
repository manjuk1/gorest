package users

type UserSerializer struct {
	user User
}

type UserResponse struct {
	ID    	 uint    `json:"id"`
	Username string  `json:"name"`
	Email    string  `json:"email"`
}

func (self *UserSerializer) Response() UserResponse {
	user := UserResponse{
		ID:		  self.user.Model.ID,
		Username: self.user.UserName,
		Email:    self.user.Email,
	}
	return user
}
package users

import (
	// "github.com/wangzitian0/golang-gin-starter-kit/common"
	// "github.com/wangzitian0/golang-gin-starter-kit/users"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/common"
	"github.com/manjuk1/gorest/db"
	"net/http"
	"fmt"
	"errors"
)

type UserApi struct {}

func (api *UserApi) createUser(c *gin.Context) {
	fmt.Println("Create User API called")
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	err := db.GetDBConn().Save(&userModelValidator.userModel).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := UserSerializer{user: userModelValidator.userModel}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func (api *UserApi) showUser(c *gin.Context) {
	fmt.Println("Show user API called")
	user := User{}
	err := db.GetDBConn().Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := UserSerializer{user: user}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})

}

func (api *UserApi) loginUser(c *gin.Context) {
	fmt.Println("User Authentication API called")
	loginValidator := LoginValidator{}
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	user := User{}
	err := db.GetDBConn().Where("email = ?", loginValidator.User.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	if user.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusUnauthorized, common.NewError("Authentication", errors.New("Unauthorized Access")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Authenticated Successfully"})

}




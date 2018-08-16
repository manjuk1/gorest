package users

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/common"
	"github.com/manjuk1/gorest/db"
	"net/http"
)

type UserApi struct{}

func (api *UserApi) CreateUser(c *gin.Context) {
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

func (api *UserApi) ShowUser(c *gin.Context) {
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

func (api *UserApi) LoginUser(c *gin.Context) {
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
	serializer := UserSerializer{user: user}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})

}

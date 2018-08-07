package users

import (
	// "github.com/wangzitian0/golang-gin-starter-kit/common"
	// "github.com/wangzitian0/golang-gin-starter-kit/users"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/common"
	"github.com/manjuk1/gorest/db"
	"net/http"
	"fmt"
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
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": &userModelValidator.userModel.ID})

}




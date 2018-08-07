package users

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/manjuk1/gorest/db"
	// "net/http"
)

func BuildRoutes(router *gin.Engine){
	api := UserApi{}
	userGrp := router.Group("/api/v1/users")
	{
		userGrp.POST("/", api.createUser)
	}

}

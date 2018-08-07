package sessions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/manjuk1/gorest/db"
	// "net/http"
)

func BuildRoutes(router *gin.Engine){
	sessionGrp := router.Group("/api/v1/sessions")
	{
		sessionGrp.POST("/", createSession)
		// sessionGrp.DELETE("/:id", deleteUser)
	}
}

func createSession(c *gin.Context){
	fmt.Println("Create Sessions API called")
}

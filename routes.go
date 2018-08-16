package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/app/users"
)

func createRoutes(router *gin.Engine) {
	userRoutes(router)
}

func userRoutes(router *gin.Engine) {
	api := users.UserApi{}
	userGrp := router.Group("/api/v1/users")
	{
		userGrp.Use(AuthenticateRequests(false))
		userGrp.POST("/", api.CreateUser)
		userGrp.POST("/login", api.LoginUser)
		userGrp.Use(AuthenticateRequests(true))
		userGrp.GET("/:id", api.ShowUser)

	}
}

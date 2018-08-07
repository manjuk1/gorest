package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/app/users"
	"github.com/manjuk1/gorest/app/sessions"
)

func createRoutes(router *gin.Engine){
  // userRoutes(router)
  // sessionsRoutes(router)	
	users.BuildRoutes(router)
	sessions.BuildRoutes(router)
}

// func userRoutes(router *gin.Engine){
// 	userGrp := router.Group("/api/v1/users")
// 	{
// 		userGrp.POST("/", users.Create)
// 		// userGrp.GET("/", listUser)
// 		// userGrp.GET("/:id", showUser)
// 		// userGrp.PUT("/:id", updateUser)
// 		// userGrp.DELETE("/:id", deleteUser)
// 	}

// }

// func sessionsRoutes(router *gin.Engine){
// 	sessionGrp := router.Group("/api/v1/sessions")
// 	{
// 		sessionGrp.POST("/", sessions.Create)
// 		// sessionGrp.DELETE("/:id", deleteUser)
// 	}
// }
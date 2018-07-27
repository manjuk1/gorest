package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
	"github.com/manjuk1/gorest/db"
)

func main() {
	fmt.Println("Starting to build a REST API")
	dbConn := db.Init()
	router := gin.Default()
	createRoutes(router)
	router.Run()
}

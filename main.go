package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/manjuk1/gorest/app/users"
	"github.com/manjuk1/gorest/config"
	"github.com/manjuk1/gorest/db"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(users.User{})
}

func main() {
	fmt.Println("Starting to build a REST API")
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}
	dbConn := db.Init()
	Migrate(dbConn)
	fmt.Errorf("DBconnected %v", dbConn.Error)
	router := gin.Default()
	createRoutes(router)
	router.Run()
}

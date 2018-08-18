package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/db"
	"github.com/jinzhu/gorm"
	"github.com/manjuk1/gorest/app/users"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(users.User{})
}

func main() {
	fmt.Println("Starting to build a REST API")
	dbConn := db.Init()
	Migrate(dbConn)
	fmt.Errorf("DBconnected %v", dbConn.Error)
	router := gin.Default()
	createRoutes(router)
	router.Run()
}

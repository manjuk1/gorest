package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
	"github.com/manjuk1/gorest/db"
	// "github.com/manjuk1/gorest/models"
	"github.com/jinzhu/gorm"
	"github.com/manjuk1/gorest/app/users"
)

func Migrate(db *gorm.DB) {
	// users.AutoMigrate()
	db.AutoMigrate(users.User{})
	// db.AutoMigrate(&articles.TagModel{})
	// db.AutoMigrate(&articles.FavoriteModel{})
	// db.AutoMigrate(&articles.ArticleUserModel{})
	// db.AutoMigrate(&articles.CommentModel{})
}


func main() {
	fmt.Println("Starting to build a REST API")
	dbConn := db.Init()
	Migrate(dbConn)
	fmt.Errorf("DBconnected %v",  dbConn.Error)
	router := gin.Default()
	createRoutes(router)
	router.Run()
}

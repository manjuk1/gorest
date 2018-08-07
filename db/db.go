package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"fmt"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=krishna dbname=golang-dev-db password=krishna port=5432")
	if err != nil {
		errStr := fmt.Errorf("Database error %s", err)
		log.Print(errStr)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

func GetDBConn() *gorm.DB {
	return DB
}
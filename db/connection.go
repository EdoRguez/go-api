package db

import (
	"log"

	"github.com/edorguez/go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var DSN = "host=localhost user=eduardo password=mysecretpassword dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DSN = "host=localhost user=postgres password=admin dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func DBConnection() {
	var error error
	DB, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil || DB == nil {
		log.Fatal(error)
	} else {
		log.Println("DB conntected to ")

		DB.AutoMigrate(models.Task{})
		DB.AutoMigrate(models.User{})
	}
}

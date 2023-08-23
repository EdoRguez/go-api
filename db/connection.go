package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DSN = "host=localhost user=eduardo password=mysecretpassword dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DSN = "host=localhost user=postgres password=admin dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// var DSN = "postgres://eduardo:mysecretpassword@localhost:5432/gorm"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB conntected to ", DB)
	}
}

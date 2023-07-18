package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){

	database, err := gorm.Open(postgres.Open("user=postgres password=postgres dbname=go_restapi_gin port=5432 sslmode=disable"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})
	DB = database
}
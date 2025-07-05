package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


var DB *gorm.DB

func ConnectDB(){
	dns := "host=localhost user=postgres password=test@123 dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err !=nil{
		panic("failed to connect to database")
	}
	fmt.Println("Connected to database successfully")
	DB = database



}
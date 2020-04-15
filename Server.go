package main

import (
	"fmt"
	"hellogin/config"
	"hellogin/models"
	"hellogin/routers"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	// creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse:", err)
	}

	defer config.DB.Close()

	// run the migrations: User
	config.DB.AutoMigrate(&models.User{})

	// server start
	r := routers.SetupRouter()
	r.Run()
}

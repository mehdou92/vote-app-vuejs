package main

import (
	"github.com/jinzhu/gorm"
	"github.com/mehdou92/vote-app-vuejs/api/config"
	"github.com/mehdou92/vote-app-vuejs/api/models"
	"github.com/mehdou92/vote-app-vuejs/api/routing"
)

var db *gorm.DB
var err error

func main() {

	// Init database
	db := config.InitDB()
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Vote{})

	router := routing.InitializeRouter()

	// Start and run the server
	router.Run(":3000")
}

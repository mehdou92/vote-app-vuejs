package main

import (
	"github.com/jinzhu/gorm"
	"github.com/lavaninho/Projet-GO/config"
	"github.com/lavaninho/Projet-GO/models"
	"github.com/lavaninho/Projet-GO/routing"
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

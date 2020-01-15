package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/lavaninho/Projet-GO/models"
	"github.com/lavaninho/Projet-GO/utils"
)

// Login login handler
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// check login payload
	var login models.Login
	err := c.BindJSON(&login)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	// get user by email
	var user models.User
	if err := db.Where("email = ?", login.Login).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// check password and set token
	if utils.CheckPassword(user, login.Password) {
		token, err := utils.GenerateJWT(user)
		if err != nil {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{"jwt": token})
	} else {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
}

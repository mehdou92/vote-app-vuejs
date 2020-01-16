package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"

	"github.com/mehdou92/vote-app-vuejs/api/models"
	"github.com/mehdou92/vote-app-vuejs/api/utils"
)

// Index index
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"OK": "User Index"})
}

// GetUser get user by id
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uuid := c.Params.ByName("id")
	var user models.User

	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers get all users
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.User

	if err := db.Find(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// InsertUser Insert user
func InsertUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	c.BindJSON(&user)

	// check validations. validate settings added to struct body.
	_, validErrs := govalidator.ValidateStruct(user)
	if validErrs != nil {
		err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	if !utils.CheckAge(user.DateOfBirth, 18) {
		err := map[string]interface{}{"error": "age limit is 18"}
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// hash user password after validate it.
	user.Password = utils.Hash(user.Password)

	// UUID
	user.UUID = uuid.Must(uuid.NewV4()).String()

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser update user by id
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uuid := c.Params.ByName("id")

	var user models.User
	var userTmp models.User

	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {

		c.BindJSON(&userTmp)

		if len(userTmp.FirstName) > 0 {
			user.FirstName = userTmp.FirstName
		}

		if len(userTmp.LastName) > 0 {
			user.LastName = userTmp.LastName
		}

		if len(userTmp.Email) > 0 {
			user.Email = userTmp.Email
		}

		if len(userTmp.Password) > 0 {
			user.Password = userTmp.Password
		}

		// check validations. validate settings added to struct body.
		_, validErrs := govalidator.ValidateStruct(user)
		if validErrs != nil {
			err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if !utils.CheckAge(user.DateOfBirth, 18) {
			err := map[string]interface{}{"error": "age limit is 18"}
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		// hash user password after validate it.
		user.Password = utils.Hash(user.Password)

		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser Delete user by id
func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uuid := c.Params.ByName("id")

	var user models.User
	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	} else {
		db.Delete(&user)
	}
}

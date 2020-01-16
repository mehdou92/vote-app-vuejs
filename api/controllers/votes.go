package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/mehdou92/vote-app-vuejs/api/models"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"OK": "Votes Index"})
}

// GetVote get vote for given id
func GetVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uuid := c.Params.ByName("id")
	var vote models.Vote

	if err := db.Where("uuid = ?", uuid).First(&vote).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

// GetVotes get all users
func GetVotes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var vote []models.Vote

	if err := db.Find(&vote).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

// InsertVote insert new vote
func InsertVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var vote models.Vote
	c.BindJSON(&vote)

	_, validErrs := govalidator.ValidateStruct(vote)
	if validErrs != nil {
		err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	vote.UUID = uuid.Must(uuid.NewV4()).String()

	if err := db.Create(&vote).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

// UpdateVote update vote for given id
func UpdateVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	claims := c.MustGet("claims").(jwt.MapClaims)

	accessLevel := int(claims["access_level"].(float64))

	uuid := c.Params.ByName("id")

	var vote models.Vote
	var voteTmp models.Vote

	if err := db.Where("uuid = ?", uuid).First(&vote).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		if accessLevel == 1 {
			c.BindJSON(&voteTmp)

			if len(voteTmp.Title) > 0 {
				vote.Title = voteTmp.Title
			}

			if len(voteTmp.Description) > 0 {
				vote.Description = voteTmp.Description
			}

			if len(voteTmp.StartDate) > 0 {
				vote.StartDate = voteTmp.StartDate
			}

			if len(voteTmp.EndDate) > 0 {
				vote.EndDate = voteTmp.EndDate
			}

			c.BindJSON(&vote)
		} else if accessLevel == 2 {
			userUUID := string(claims["uuid"].(string))

			if vote.UUIDVote != nil {

				for _, voter := range vote.UUIDVote {
					if voter == userUUID {
						err := map[string]interface{}{"error": "The proposal was already voted"}
						c.JSON(http.StatusUnauthorized, err)
						return
					}
				}

				vote.UUIDVote = append(vote.UUIDVote, userUUID)

			} else {

				var newUUIDVotes []string
				vote.UUIDVote = append(newUUIDVotes, userUUID)
			}
		}

		// check validations. validate settings added to struct body.
		_, validErrs := govalidator.ValidateStruct(vote)
		if validErrs != nil {
			err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		db.Save(&vote)
		c.JSON(http.StatusOK, vote)
	}
}

// DeleteVote delete vote for given id
func DeleteVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	uuid := c.Params.ByName("id")

	var vote models.Vote
	if err := db.Where("uuid = ?", uuid).First(&vote).Error; err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	} else {
		db.Delete(&vote)
	}
}

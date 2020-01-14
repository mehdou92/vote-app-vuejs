package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lavaninho/Projet-GO/utils"
)

// AccessLevel type
type AccessLevel int

const (
	// Admin admin access level constant
	Admin AccessLevel = 1
	// User user access level constant
	User AccessLevel = 2
)

// Authorization check user authorization
func Authorization(needAccessLevel bool, accessLevels ...AccessLevel) gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract authorization header
		headerString := c.GetHeader("Authorization")

		// if not, throw error
		if len(headerString) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Not authorized.",
			})

			c.Abort()
			return
			//if token format not valid, throw error.
		} else if headerPairs := strings.Split(headerString, " "); len(headerPairs) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header format is invalid.",
			})
			c.Abort()
			return
			// if jwt token does not exist, throw error.
		} else if claims, err := utils.CheckJwtToken(headerPairs[1]); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		} else if needAccessLevel {
			slices := make([]int, len(accessLevels))
			for i := range accessLevels {
				al := accessLevels[i]
				value := int(al)
				slices = append(slices, value)
			}
			if _, exists := utils.FindAccessLevel(slices, int(claims["access_level"].(float64))); !exists {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Not Authorized",
				})
				c.Abort()
				return
			}
		} else {
			c.Set("claims", claims)
			c.Next()
		}
	}
}

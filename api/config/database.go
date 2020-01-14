package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database structure
type Database struct {
	*gorm.DB
}

// DB database instance
var DB *gorm.DB
var err error

// InitDB opens a database and saves the reference to `Database` struct.
func InitDB() *gorm.DB {
	var db = DB

	driver := "postgres"
	database := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	if driver == "postgres" {

		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}

	db.LogMode(true)
	DB = db

	return DB
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// Inject set db variable in gin
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

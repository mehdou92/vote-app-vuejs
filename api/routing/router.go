package routing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavaninho/Projet-GO/config"
	"github.com/lavaninho/Projet-GO/controllers"
	"github.com/lavaninho/Projet-GO/middlewares"
)

// InitializeRouter initialize router
func InitializeRouter() *gin.Engine {

	router := gin.Default()
	router.ForwardedByClientIP = true

	db := config.GetDB()

	if db == nil {
		log.Fatal("DB Error")
	}

	router.Use(config.Inject(db))

	// Use middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Use custom middlewares
	router.Use(middlewares.CORS())

	api := router.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// public endpoints.

	loginLimiter, ok := middlewares.LoginLimiter()

	if !ok {
		log.Fatal("Redis Error")
	}

	login := api.Group("/login")
	login.Use(loginLimiter)
	{
		login.POST("/", controllers.Login)
	}

	// User routes
	users := api.Group("/users")

	users.Use(middlewares.Authorization(true, 1))
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.POST("/", controllers.InsertUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	// Vote routes
	votesAdmin := api.Group("/votes")

	votesAdmin.Use(middlewares.Authorization(true, 1))
	{
		votesAdmin.POST("/", controllers.InsertVote)
		votesAdmin.DELETE("/:id", controllers.DeleteVote)
	}

	votes := api.Group("/votes")

	votes.Use(middlewares.Authorization(false))
	{
		votes.GET("/", controllers.GetVotes)
		votes.GET("/:id", controllers.GetVote)
		votes.PUT("/:id", controllers.UpdateVote)
	}

	return router
}

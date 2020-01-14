package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"

	"github.com/lavaninho/Projet-GO/config"
)

// LoginLimiter Login rate limiter middleware
func LoginLimiter() (gin.HandlerFunc, bool) {

	var middleware gin.HandlerFunc

	store, rate, ok := config.InitRedis()

	if !ok {
		return middleware, false
	}

	middleware = mgin.NewMiddleware(limiter.New(store, rate))

	return middleware, true
}

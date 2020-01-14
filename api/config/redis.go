package config

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/ulule/limiter/v3"

	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var rate limiter.Rate
var store limiter.Store
var option *redis.Options
var client *redis.Client

// InitRedis init redis with login limiter
func InitRedis() (limiter.Store, limiter.Rate, bool) {
	// Define a limit rate to 3 requests per minute.
	rate, err = limiter.NewRateFromFormatted("3-M")
	if err != nil {
		log.Fatal(err)
		return store, rate, false
	}

	// Create a redis client.
	option, err = redis.ParseURL("redis://redis:6379/0")
	if err != nil {
		log.Fatal(err)
		return store, rate, false
	}
	client = redis.NewClient(option)

	// Create a store with the redis client.
	store, err = sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter-login",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		return store, rate, false
	}

	return store, rate, true
}

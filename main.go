package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

const requestsPerMinute = 3

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})

	// API
	router := gin.Default()
	router.GET("/limit/:id", checkLimit)
	err := router.Run(":1111")
	if err != nil {
		return
	}
}

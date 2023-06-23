package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

// checkLimit of user requests using sliding window algorithm
func checkLimit(ginContext *gin.Context) {
	userId := ginContext.Param("id")

	currentTime := time.Now()
	timeToCompare := currentTime.Add(-time.Minute * 1).Unix()

	// Remove all the timestamps from the Sorted Set that are older than "CurrentTime - 1 minute".
	_, err := redisClient.ZRemRangeByScore(context.Background(), userId, "0", fmt.Sprintf("(%d", timeToCompare)).Result()
	if err != nil {
		fmt.Println("Error while removing timestamps: ", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Count the total number of elements in the sorted set. Reject the request if this count is greater than our throttling limit.
	count, err := redisClient.ZCard(context.Background(), userId).Result()
	if err != nil {
		fmt.Println("Error on counting: ", err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	if count > requestsPerMinute-1 {
		fmt.Println("request rejected")
		ginContext.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Request was rejected - your limit reached!"})
		return
	}

	// Insert the current time in the sorted set and accept the request.
	redisClient.ZAdd(context.Background(), userId, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: time.Now().Unix(),
	})
}

package main

import (
	// Redis service

	"fmt"
	"one-time-secret/goredis"
	"time"

	"github.com/google/uuid"

	// Service level modules

	// Gin server for rest api
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type SECRET struct {
	ENCRYPT_KEY string `json:"encryptKey" binding:"required"`
	SECRET      string `json:"secret" binding:"required"`
}

type GET_SECRET_REQUEST struct {
	REQUEST_ID string `json:"requestId" binding:"required"`
}

func main() {

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	goredis.InitPool()

	// Create secret with key and value in redis
	api.POST("/v1/secret/create", func(c *gin.Context) {
		var secret SECRET
		c.BindJSON(&secret)
		requestId := uuid.New().String()

		key, value, err := goredis.Set(requestId, secret.SECRET)
		if err == nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"requestId":    requestId,
				"key":          key,
				"secret":       value,
				"responseTime": time.Now(),
			})
		} else {
			c.JSON(400, gin.H{
				"message": "failed to save record.",
			})
		}
	})

	// Get secret with key from redis
	api.GET("/v1/secret/get", func(c *gin.Context) {
		var request GET_SECRET_REQUEST
		c.BindJSON(&request)
		fmt.Print("Request id : %s", request.REQUEST_ID)

		key, value, err := goredis.Get(request.REQUEST_ID)

		fmt.Println("Values from Redis  Key : %s , Value : %s", key, value)
		if err == nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"secret":       value,
				"responseTime": time.Now(),
			})
		} else {
			c.JSON(400, gin.H{
				"message":      "Failed to get key :" + key,
				"responseTime": time.Now(),
			})
		}

	})

	// Simple monitor API to check if service is up and running.
	api.GET("/monitor", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status":       "READY",
			"responseTime": time.Now(),
		})
	})

	// Specifying on which port to run
	r.Run(":8081")
}

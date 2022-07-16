package main

import (
	// Redis service

	"fmt"
	"one-time-secret/goredis"

	// Service level modules

	// Gin server for rest api
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type SECRET struct {
	KEY   string `json:"key" binding:"required"`
	VALUE string `json:"value" binding:"required"`
}

func main() {

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	goredis.InitPool()

	// Create secret with key and value in redis
	api.POST("/secret/v1/set", func(c *gin.Context) {
		var secret SECRET
		c.BindJSON(&secret)
		key, value, err := goredis.Set(secret.KEY, secret.VALUE)
		if err == nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"key":   key,
				"value": value,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "failed to save record.",
			})
		}
	})

	// Get secret with key from redis
	api.POST("/secret/v1/get", func(c *gin.Context) {
		var secret SECRET
		c.BindJSON(&secret)
		key, value, err := goredis.Get(secret.KEY)
		if err == nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"key":   key,
				"value": value,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "failed to get record.",
			})
		}
	})

	// Simple monitor API to check if API is working
	api.GET("/monitor", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "READY",
		})
	})

	// Specifying on which port to run
	r.Run(":8081")
}

package main

import (
	// Redis service

	"one-time-secret/goredis"

	// Service level modules

	// Gin server for rest api
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

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
	api.POST("/createSecret", func(c *gin.Context) {
		var secret SECRET
		c.BindJSON(&secret)
		key, value, err := goredis.Set(secret.KEY, secret.VALUE)
		if err != nil {
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

	api.POST("/testPost", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.USER}) // Your custom response here
	})

	// Simple monitor API to check if API is working
	api.GET("/monitor", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "READY",
		})
	})

	// Specifying on which post to run
	r.Run(":8080")
}

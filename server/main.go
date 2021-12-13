package main

import (
	"fmt"

	// Service level modules
	"one-time-secret/service"

	// Gin server for rest api
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")

	//name := service.Test()
	//fmt.Println("Name is : ", name)

	api.GET("/ping", func(c *gin.Context) {
		name := service.Test()
		fmt.Println("Name is : ", name)
		c.JSON(200, gin.H{
			"message": name,
		})
	})

	api.POST("/ping", func(c *gin.Context) {
		name := service.Test()
		fmt.Println("Name is : ", name)
		c.JSON(200, gin.H{
			"message": name,
		})
	})

	api.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "READY",
		})
	})

	// Specifying on which post to run
	r.Run(":8080")
}

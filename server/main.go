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

func main() {

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	goredis.InitPool()
	//name := service.Test()
	//fmt.Println("Name is : ", name)
	// api.GET("/ping", func(c *gin.Context) {
	// 	name := service.Test()
	// 	fmt.Println("Name is : ", name)
	// 	c.JSON(200, gin.H{
	// 		"message": name,
	// 	})
	// })

	api.POST("/post", func(c *gin.Context) {
		var login LOGIN
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.USER}) // Your custom response here
	})

	api.GET("/monitor", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "READY",
		})
	})

	// Specifying on which post to run
	r.Run(":8080")
}

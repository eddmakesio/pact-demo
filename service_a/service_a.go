package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/bing", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"bing": "bong",
		})
	})
	r.Run(":10001")
}
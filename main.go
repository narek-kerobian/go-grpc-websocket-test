package main

import (
	"github.com/gin-gonic/gin"
)

const appPort = "10080"

func main() {
    r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + appPort)
}

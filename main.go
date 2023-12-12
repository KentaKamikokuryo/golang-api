package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kentakamikokuryo/golang-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!!!!",
		})
	})
	r.Run()
}

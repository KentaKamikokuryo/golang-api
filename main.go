package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kentakamikokuryo/golang-api/controllers"
	"github.com/kentakamikokuryo/golang-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)       // Create a new post
	r.GET("/posts", controllers.PostsIndex)         // Road all posts
	r.GET("/posts/:id", controllers.PostsShow)      // Road a post by id
	r.PUT("/posts/:id", controllers.PostsUpdate)    // Update a post by id
	r.DELETE("/posts/:id", controllers.PostsDelete) // Delete a post by id
	r.Run()
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kentakamikokuryo/golang-api/initializers"
	"github.com/kentakamikokuryo/golang-api/models"
)

func PostsCreate(c *gin.Context) {

	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Create new post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return new post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Find post by id
	var post models.Post
	initializers.DB.First(&post, id)

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find post by id
	var post models.Post
	initializers.DB.First(&post, id)

	// Update post
	initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		},
	)

	// Return updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Delete post
	initializers.DB.Delete(&models.Post{}, id)

	// Return deleted post
	c.JSON(200, gin.H{
		"delete": "delete!!!",
	})
}

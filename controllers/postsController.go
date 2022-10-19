package controllers

import (
	"github.com/amos101/go-crud/initializers"
	"github.com/amos101/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// creat a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Postsindex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post where updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)

}
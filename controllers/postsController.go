package controllers

import (
	"first_go_api/initializers"
	"first_go_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data of req body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create s post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGetAll(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostGetById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	//get data from body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find
	var post models.Post
	initializers.DB.First(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//Response

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Find
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	//Delete
	initializers.DB.Delete(&post)

	//Response
	c.Status(204)
}

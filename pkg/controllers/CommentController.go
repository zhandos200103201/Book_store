package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var body struct {
		Title   string
		Comment string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title or comment is wrong",
		})
		return
	}

	comment := models.Comment{Author: GetUserEmail(c), Title: body.Comment, Book: body.Title}
	result := initializers.GetDB().Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create new comment",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"New comment": comment,
	})
}

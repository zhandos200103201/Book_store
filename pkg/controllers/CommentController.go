package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// comment system
// for creating a new comment

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

	var user models.User
	initializers.GetDB().Find(&user, "email=?", GetUserEmail(c)) // getting author of comment by "GetUserEmail"

	comment := models.Comment{Author: GetUserEmail(c), Title: body.Comment, Book: body.Title} // create a new comment
	result := initializers.GetDB().Create(&comment)                                           // insert it into table
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create new comment",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"New comment": comment,
	})
}

// for removing the comment by id

func DeleteComment(c *gin.Context) {
	id := c.Param("id") // get target comment id
	var targetComment models.Comment
	initializers.GetDB().Find(&targetComment, "id=?", id) // getting comment with by id
	var admin models.User
	initializers.GetDB().Find(&admin, "email=?", GetUserEmail(c)) // getting user who want to delete the comment

	if GetUserEmail(c) != targetComment.Author && admin.Type != "Admin" { // checking for owners the comment
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Only owners can delete his comments",
		})
		return
	}

	initializers.GetDB().Delete(&targetComment, "id=?", id) // if its comment owner then we delete it
	c.JSON(http.StatusOK, gin.H{
		"Comment": "Successfully removed",
	})
}

// for getting all comments from table

func GetAllComments(c *gin.Context) {
	var comments []models.Comment
	initializers.GetDB().Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"Comments": comments,
	})
}

// for getting target comments for books (by id)

func GetCommentsByID(c *gin.Context) {
	id := c.Param("id") // get target id
	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id) // take target book
	var comments []models.Comment
	initializers.GetDB().Find(&comments, "book=?", target.Title) // get all comment for this book
	c.JSON(http.StatusOK, gin.H{
		"Comments": comments,
	})
}

// for getting target comments for books (by book titles)

func GetCommentsForBook(title string) []models.Comment {
	var comments []models.Comment
	initializers.GetDB().Find(&comments, "book=?", title)
	return comments
}

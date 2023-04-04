package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var books []models.Book

func GetBooks(c *gin.Context) {
	initializers.GetDB().Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"Books": books,
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"Book": target,
	})
}

func CreateBook(c *gin.Context) {
	book := models.Book{}
	if c.Bind(&book) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	result := initializers.GetDB().Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a new book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func EditBook(c *gin.Context) {
	id := c.Param("id")
	book := models.Book{}
	if c.Bind(&book) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id)
	if book.Title != target.Title {
		initializers.GetDB().Model(&target).Update("Title", book.Title)
	}
	if book.Description != target.Description {
		initializers.GetDB().Model(&target).Update("Description", book.Description)
	}
	if book.Author != target.Author {
		initializers.GetDB().Model(&target).Update("Author", book.Author)
	}
	if book.Rating != target.Rating {
		initializers.GetDB().Model(&target).Update("Rating", book.Rating)
	}
	if book.Price != target.Price {
		initializers.GetDB().Model(&target).Update("Price", book.Price)
	}
	if book.Language != target.Language {
		initializers.GetDB().Model(&target).Update("Language", book.Language)
	}
	if book.Image != target.Image {
		initializers.GetDB().Model(&target).Update("Image", book.Image)
	}
	if book.Category != target.Category {
		initializers.GetDB().Model(&target).Update("Category", book.Category)
	}
	c.JSON(http.StatusOK, gin.H{
		"Edit": "Book successfully edited",
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var target models.Book
	initializers.GetDB().Delete(&target, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"Book": "removed",
	})
}

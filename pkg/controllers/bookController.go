package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// for getting all books from db table
func GetBooks(c *gin.Context) {
	var books []models.Book
	initializers.GetDB().Find(&books)
	//c.HTML(http.StatusOK, "books.html", gin.H{
	//	"Books": books,
	//})
	c.JSON(http.StatusOK, gin.H{
		"Books": books,
	})
}

// for getting target book
func GetBook(c *gin.Context) {
	id := c.Param("id") // get target book id from browser params
	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id) // get book by with target id
	comments := GetCommentsForBook(target.Title)   // getting comments for this book
	c.HTML(http.StatusOK, "book.html", gin.H{      // give all data for book.html file
		"Book":     target,
		"Comments": comments,
	})
}

// creating a new book and adding it to database
func CreateBook(c *gin.Context) {
	book := models.Book{}
	if c.Bind(&book) != nil { // getting book fields from site, then assign it to created book model
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	result := initializers.GetDB().Create(&book) // inserting book into table
	if result.Error != nil {                     // if book is exists db return the error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a new book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{}) // else we return empty json file
}

// for editing the book

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
	initializers.GetDB().Find(&target, "id=?", id) // getting target book from table
	// then replace it compare by each field

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

// for removing the book

func DeleteBook(c *gin.Context) {
	id := c.Param("id") // taking target book id from site
	var target models.Book
	initializers.GetDB().Delete(&target, "id=?", id) // take the target book with id, then delete it
	c.JSON(http.StatusOK, gin.H{
		"Book": "removed",
	})
}

// filtering the books by arrange two price

func PriceFiltering(c *gin.Context) {
	from := c.Query("from") // getting first price
	to := c.Query("to")     // getting second price
	var books []models.Book
	// get all books from table with price between this two prices
	result := initializers.GetDB().Where("price >= ?", from).Where("price <= ?", to).Order("price").Find(&books)
	// return the error if failed to get books between two price
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get books between of prices",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

// filtering books by rating

func RatingFiltering(c *gin.Context) {
	rating := c.Query("rating") // getting book rating
	var books []models.Book     // taking all books with rating >= target rating
	result := initializers.GetDB().Where("rating >= ?", rating).Order("rating desc").Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get books by rating",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}
func Search(c *gin.Context) {
	name := c.Query("name") + "%" // getting book title
	var books []models.Book
	result := initializers.GetDB().Where("title like ?", name).Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get books by title",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

// for giving a ratings by clients

func GiveRating(c *gin.Context) {
	var body struct { // taking book title and rating for the book
		Title  string //target title
		Rating float32
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var target models.Book
	initializers.GetDB().Find(&target, "title=?", body.Title) // taking book with target title
	avg := (body.Rating + target.Rating) / 2
	if body.Rating <= 5.0 { // if giving rating less or equal to 5 point then we update the rating for this book
		initializers.GetDB().Model(&target).Update("Rating", avg)
	}
	c.JSON(http.StatusOK, gin.H{
		"Change book rating": avg,
	})
}

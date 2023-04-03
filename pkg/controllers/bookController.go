package controllers

import (
	"Test2/pkg/models"
	"Test2/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var books []models.Book

func GetBooks(w http.ResponseWriter, _ *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	Id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("Error during the parsing")
	}
	targetBook, _ := models.GetBookById(int16(Id))
	res, _ := json.Marshal(targetBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := utils.ParseBody(r)
	book := newBook.CreateBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	editBook := utils.ParseBody(r)
	params := mux.Vars(r)
	bookId := params["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error during the parsing")
	}
	book, db := models.GetBookById(int16(Id))
	book.Title = editBook.Title
	book.Description = editBook.Description
	book.Author = editBook.Author
	book.Category = editBook.Category
	book.Image = editBook.Image
	book.Language = editBook.Language
	book.Rating = editBook.Rating
	book.Price = editBook.Price
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error during the parsing")
	}
	targetBook := models.DeleteBook(int16(Id))
	res, _ := json.Marshal(targetBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

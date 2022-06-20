package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func main() {
	r := gin.New()

	// GET /books
	r.GET("/books", getBooks)

	// GET /books/:id
	r.GET("/books/:id", getBook)

	// POST /books
	r.POST("/books", createBook)

	// PUT /books/:id
	r.PUT("/books/:id", updateBook)

	// DELETE /books/:id
	r.DELETE("/books/:id", deleteBook)

	r.Run(":5000")
}

func updateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + err.Error(),
		})
		return
	}

	updatedBook := Book{
		ID:     id,
		Title:  book.Title,
		Author: book.Author,
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func getBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book Book

	for i, b := range books {
		if b.ID == id {
			book = books[i]
		}
	}

	if book.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Error: Book not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func getBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, books)
}

func createBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + err.Error(),
		})
		return
	}

	books = append(books, book)

	ctx.JSON(http.StatusCreated, book)
}

func deleteBook(ctx *gin.Context) {

	id := ctx.Param("id")

	var book Book

	for i, b := range books {
		if b.ID == id {
			book = books[i]
		}
	}

	if book.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Error: Book not found",
		})
		return
	}

	// นำออกจาก slice
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	ctx.Status(http.StatusNoContent)
}

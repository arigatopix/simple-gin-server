package main

import (
	"fmt"
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

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "world",
		})
	})

	// GET /books
	r.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, books)
	})

	// POST /books
	r.POST("/books", func(ctx *gin.Context) {
		var book Book

		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Error:" + err.Error(),
			})
			return
		}

		books = append(books, book)

		fmt.Println(books)

		ctx.JSON(http.StatusCreated, book)
	})

	// DELETE /books/:id
	r.DELETE("/books/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		// นำออกจาก slice
		for i, book := range books {
			if book.ID == id {
				books = append(books[:i], books[i+1:]...)
			}
		}

		ctx.Status(http.StatusNoContent)
	})

	r.Run(":5000")
}

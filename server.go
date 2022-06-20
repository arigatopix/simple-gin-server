package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *gorm.DB

func main() {
	var err error

	db, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	db.DB()

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

	updateBook := Book{
		ID:     id,
		Title:  book.Title,
		Author: book.Author,
	}

	if result := db.Model(&Book{}).Where("ID = ?", id).Updates(updateBook); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &updateBook)
}

func getBook(ctx *gin.Context) {

	id := ctx.Param("id")

	var book Book

	if result := db.First(&book, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

func getBooks(ctx *gin.Context) {
	var books []Book

	if result := db.Find(&books); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &books)
}

func createBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + err.Error(),
		})
		return
	}

	result := db.Create(&book)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, &book)
}

func deleteBook(ctx *gin.Context) {

	id := ctx.Param("id")

	if result := db.First(&Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	if result := db.Delete(&Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

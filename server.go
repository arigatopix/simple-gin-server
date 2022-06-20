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

type Handler struct {
	db *gorm.DB
}

func connectDB(db *gorm.DB) *Handler {
	return &Handler{db}
}

func main() {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	// create new connection instance
	// แต่ละ route เป็น instance ของ db
	handler := connectDB(db)

	r := gin.New()

	// GET /books
	r.GET("/books", handler.getBooks)

	// GET /books/:id
	r.GET("/books/:id", handler.getBook)

	// POST /books
	r.POST("/books", handler.createBook)

	// PUT /books/:id
	r.PUT("/books/:id", handler.updateBook)

	// DELETE /books/:id
	r.DELETE("/books/:id", handler.deleteBook)

	r.Run(":5000")
}

func (h *Handler) updateBook(ctx *gin.Context) {
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

	// h.db -> เรียกใช้ struct Handler.db นะได้ db *gorm.DB
	// เรียก instance ใหม่ทุกครั้งที่มี route เข้ามา
	if result := h.db.Model(&Book{}).Where("ID = ?", id).Updates(updateBook); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &updateBook)
}

func (h *Handler) getBook(ctx *gin.Context) {

	id := ctx.Param("id")

	var book Book

	if result := h.db.First(&book, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

func (h *Handler) getBooks(ctx *gin.Context) {
	var books []Book

	if result := h.db.Find(&books); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &books)
}

func (h *Handler) createBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + err.Error(),
		})
		return
	}

	result := h.db.Create(&book)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, &book)
}

func (h *Handler) deleteBook(ctx *gin.Context) {

	id := ctx.Param("id")

	if result := h.db.First(&Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	if result := h.db.Delete(&Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

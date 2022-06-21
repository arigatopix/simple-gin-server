package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "gin-webservice/db"
	model "gin-webservice/models"
)

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	db := db.ConnectDB()

	var book model.Book

	if result := db.First(&book, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	var bookFromBody model.Book

	if err := ctx.ShouldBindJSON(&bookFromBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + err.Error(),
		})
		return
	}

	updatedBook := model.Book{
		ID:     book.ID,
		Title:  bookFromBody.Title,
		Author: bookFromBody.Author,
	}

	// h.db -> เรียกใช้ struct Handler.db นะได้ db *gorm.DB
	// เรียก instance ใหม่ทุกครั้งที่มี route เข้ามา
	if result := db.Model(&model.Book{}).Where("ID = ?", id).Updates(updatedBook); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &updatedBook)
}

func GetBook(ctx *gin.Context) {
	db := db.ConnectDB()

	id := ctx.Param("id")

	var book model.Book

	if result := db.First(&book, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

func GetBooks(ctx *gin.Context) {
	db := db.ConnectDB()

	var books []model.Book

	if result := db.Find(&books); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &books)
}

func CreateBook(ctx *gin.Context) {
	db := db.ConnectDB()

	var book model.Book

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

func DeleteBook(ctx *gin.Context) {
	db := db.ConnectDB()

	id := ctx.Param("id")

	if result := db.First(&model.Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	if result := db.Delete(&model.Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error:" + result.Error.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

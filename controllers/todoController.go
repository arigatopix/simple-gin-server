package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Todo",
	})
}

func GetTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Todos",
	})
}

func CreateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create Todo",
	})
}

func UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Todo",
	})
}

func DeleteTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Todo",
	})
}

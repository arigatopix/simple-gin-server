package main

import (
	"github.com/gin-gonic/gin"

	c "gin-webservice/controllers"
)

func main() {
	r := gin.New()

	// GET /books
	r.GET("/books", c.GetBooks)

	// GET /books/:id
	r.GET("/books/:id", c.GetBook)

	// POST /books
	r.POST("/books", c.CreateBook)

	// PUT /books/:id
	r.PUT("/books/:id", c.UpdateBook)

	// DELETE /books/:id
	r.DELETE("/books/:id", c.DeleteBook)

	r.Run(":5000")
}

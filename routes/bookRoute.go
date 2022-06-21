package book

import (
	c "gin-webservice/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup) {
	route := group.Group("/books")
	{
		// GET /books
		route.GET("/", c.GetBooks)

		// GET /books/:id
		route.GET("/:id", c.GetBook)

		// POST /books
		route.POST("/", c.CreateBook)

		// PUT /books/:id
		route.PUT("/:id", c.UpdateBook)

		// DELETE /books/:id
		route.DELETE("/:id", c.DeleteBook)
	}
}

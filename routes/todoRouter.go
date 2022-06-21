package routes

import (
	c "gin-webservice/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(group *gin.RouterGroup) {
	route := group.Group("/todos")
	{
		// GET /todos
		route.GET("/", c.GetTodos)

		// GET /todos/:id
		route.GET("/:id", c.GetTodo)

		// POST /todos
		route.POST("/", c.CreateTodo)

		// PUT /todos/:id
		route.PUT("/:id", c.UpdateTodo)

		// DELETE /todos/:id
		route.DELETE("/:id", c.DeleteTodo)
	}
}

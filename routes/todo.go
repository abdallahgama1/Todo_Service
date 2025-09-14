package routes

import (
	"Todo_Service/handlers"
	"Todo_Service/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine, todoHandler *handlers.TodoHandler) {
	todoRoutes := router.Group("/todos")

	todoRoutes.Use(middlewares.AuthMiddleware())
	{
		todoRoutes.POST("", todoHandler.CreateTodo)
		todoRoutes.GET("", todoHandler.GetAllTodos)

		todoRoutes.GET("/:id", todoHandler.GetTodoByID)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
		
		todoRoutes.GET("/category/:category", todoHandler.GetTodosByCategory)
		todoRoutes.PUT("/category/:category", todoHandler.UpdateStatusByCategory)
		
		todoRoutes.GET("/status/:status", todoHandler.GetTodosByStatus)
		todoRoutes.GET("/search", todoHandler.SearchTodosByTitle)

		todoRoutes.DELETE("", todoHandler.DeleteAllTodos)
	}
}
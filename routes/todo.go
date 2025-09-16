package routes

import (
	"Todo_Service/handlers"
	middlewares "Todo_Service/middleware"

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

	}
}

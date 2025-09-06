package routes

import (
	"Todo_Service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.TodoHandler) *gin.Engine {
	router := gin.Default()

	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("", h.GetAllTodos)
		todoRoutes.GET("/:id", h.GetTodoByID)
		todoRoutes.GET("/category/:category", h.GetTodosByCategory)
		todoRoutes.GET("/status/:status", h.GetTodosByStatus)
		todoRoutes.GET("/search", h.SearchTodosByTitle)
		todoRoutes.POST("", h.CreateTodo)
		todoRoutes.PUT("/:id", h.UpdateTodo)
		todoRoutes.PUT("/category/:category", h.UpdateStatusByCategory)
		todoRoutes.DELETE("/:id", h.DeleteTodo)
		todoRoutes.DELETE("", h.DeleteAllTodos)
	}

	return router
}

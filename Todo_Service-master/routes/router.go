package routes

import (
	"Todo_Service/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler) *gin.Engine {
	router := gin.Default()

	RegisterUserRoutes(router, userHandler)
	RegisterTodoRoutes(router, todoHandler)

	return router
}
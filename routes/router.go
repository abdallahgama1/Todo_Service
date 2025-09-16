package routes

import (
	_ "Todo_Service/docs"
	"Todo_Service/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterUserRoutes(router, userHandler)
	RegisterTodoRoutes(router, todoHandler)

	return router
}

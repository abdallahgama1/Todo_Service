package handlers_test

import (
	"Todo_Service/handlers"
	"Todo_Service/models"
	"Todo_Service/repositories"
	"Todo_Service/routes"
	"Todo_Service/services"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testRouter *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	testRouter = setupTestRouter()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupTestRouter() *gin.Engine {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to real test database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	return routes.SetupRouter(userHandler, todoHandler)
}

package main

import (
	"Todo_Service/config"
	"Todo_Service/handlers"
	"Todo_Service/models"
	"Todo_Service/repositories"
	"Todo_Service/routes"
	"Todo_Service/services"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	log.Println("Database connection successful.")

	log.Println("Running database migrations...")
	db.AutoMigrate(&models.Todo{})

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	router := routes.SetupRouter(todoHandler)
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

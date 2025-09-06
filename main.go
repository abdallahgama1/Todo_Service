package main

import (
	"Todo_Service/config"
	"Todo_Service/global"
	"Todo_Service/handlers"
	"Todo_Service/repositories"
	"Todo_Service/routes"
	"Todo_Service/services"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping the database: %v", err)
	}
	log.Println("Database connection successful.")
	global.DB = db

	todoRepo := repositories.NewTodoRepository(global.DB)
	if err := todoRepo.InitDB(); err != nil {
		log.Fatalf("could not initialize database table: %v", err)
	}

	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	router := routes.SetupRouter(todoHandler)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

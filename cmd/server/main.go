package main

import (
	"ainyx-backend/config"
	db "ainyx-backend/db/sqlc"
	"ainyx-backend/internal/handler"
	"ainyx-backend/internal/logger"
	"ainyx-backend/internal/middleware"
	"ainyx-backend/internal/repository"
	"ainyx-backend/internal/routes"
	"ainyx-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Initialize logger
	logger.Init()
	defer logger.Log.Sync()

	// 2. Connect to database
	cfg := config.Config{
		DBHost:     "localhost",
		DBPort:     "5433",
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "ainyx_db",
	}
	database := config.Connect(cfg)
	defer database.Close()

	// 3. Initialize SQLC queries
	queries := db.New(database)

	// 4. Initialize layers
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 5. Setup Fiber app
	app := fiber.New()

	// 6. Register middleware
	app.Use(middleware.RequestLogger)

	// 7. Setup routes
	routes.SetupRoutes(app, userHandler)

	// 8. Start server
	logger.Log.Info("Server starting on port 3000")
	if err := app.Listen(":3000"); err != nil {
		logger.Log.Fatal("Failed to start server")
	}
}
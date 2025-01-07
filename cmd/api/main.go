package main

import (
	"fmt"

	"github.com/channyeintun/teacher-booking-system/config"
	"github.com/channyeintun/teacher-booking-system/db"
	"github.com/channyeintun/teacher-booking-system/handlers"
	"github.com/channyeintun/teacher-booking-system/repositories"
	"github.com/channyeintun/teacher-booking-system/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Teacher Booking System",
		ServerHeader: "Fiber",
	})

	// Repositories
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	// privateRoutes := server.Use(middlewares.AuthProtected(db))

	// handlers.NewEventHandler(privateRoutes.Group("/protected"), someRepository)

	app.Listen(fmt.Sprintf("%s", ":" + envConfig.ServerPort))
}

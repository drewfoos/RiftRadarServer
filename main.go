package main

import (
	"os"

	"github.com/drewfoos/RiftRadarServer/internal/handlers"
	"github.com/drewfoos/RiftRadarServer/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	handlers.InitApiKey()

	// Create a new Fiber instance
	app := fiber.New()

	// Add middleware
	app.Use(middleware.Cors())

	// Define routes
	app.Post("/search", handlers.SearchHandler)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}

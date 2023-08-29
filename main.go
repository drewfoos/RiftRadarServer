package main

import (
	"fmt"

	"github.com/drewfoos/RiftRadarServer/vendor/handlers"
	"github.com/drewfoos/RiftRadarServer/vendor/middleware"

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

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	env := os.Getenv("ENV")
	allowOrigins := "https://riftradar.vercel.app"
	if env == "dev" {
		allowOrigins = "http://localhost:3000"
	}
	return cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowHeaders: "Origin, Content-Type",
		AllowMethods: "POST, OPTIONS",
	})
}

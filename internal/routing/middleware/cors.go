package middleware

import (
	"github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCorsPolicy(allowedOrigins, allowedMethods, allowedHeaders string) fiber.Handler {
	corsPolicy := cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: allowedMethods,
		AllowHeaders: allowedHeaders,
	})
	return corsPolicy
}

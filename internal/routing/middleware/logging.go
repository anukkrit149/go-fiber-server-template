package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func NewCustomLogger() fiber.Handler {
	loggerConfig := logger.New(logger.Config{
		Format:     "${pid} ${locals:requestid} ${status} - ${method} ${path}",
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "Asia/Kolkata",
		Output:     os.Stderr,
	})
	return loggerConfig
}

package routing

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-webserver-template/internal/config"
)

func Initialize(app *fiber.App, config config.Core) {
	api := app.Group("/api")
	//api.Use()
	//api.Use() todo: logging Middleware
	//api.Use() todo: prometheus Middleware
	// todo: cors policy
	//
	api.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"status": "success",
		})
	})

	app.Get("/commit.txt", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"commit": config.GitCommitHash,
		})
	})

}

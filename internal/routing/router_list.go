package routing

import "github.com/gofiber/fiber/v2"

type Route struct {
	Group      string
	Middleware []handlerFunc
	Endpoints  []endpoint
}

type endpoint struct {
	Method  string
	Path    string
	Handler handlerFunc
}

type handlerFunc func(ctx *fiber.Ctx) error

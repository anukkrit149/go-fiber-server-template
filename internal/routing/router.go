package routing

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"go-rest-webserver-template/internal/config"
	v1 "go-rest-webserver-template/internal/controller/v1"
	"go-rest-webserver-template/internal/routing/middleware"
)

func Initialize(app *fiber.App, config config.Core) {

	// Prometheus Integration
	prometheus := fiberprometheus.New(config.ServiceName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Base Route Grp - API
	api := app.Group("/api")
	//Basic Middleware for app
	api.Use(middleware.NewCorsPolicy(allowedOrigins, allowedMethods, allowedHeaders))
	api.Use(middleware.NewCustomLogger())

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

func registerRouterGroup(grp RouteGrp, app *fiber.App) {
	routeGrp := app.Group(grp.Group)

	// registering all middlewares
	for _, middlewareHandler := range grp.Middleware {
		routeGrp.Use(middlewareHandler)
	}
	// registering endpoints
	for _, route := range grp.Endpoints {
		switch route.Method {
		case fiber.MethodPost:
			routeGrp.Post(route.Path, route.Handler)
			break
		case fiber.MethodGet:
			routeGrp.Put(route.Path, route.Handler)
			break
		case fiber.MethodPatch:
			routeGrp.Patch(route.Path, route.Handler)
			break
		case fiber.MethodPut:
			routeGrp.Put(route.Path, route.Handler)
			break
		}
	}
}

func getRouterList() []RouteGrp {
	var routes = []RouteGrp{
		{
			Group:      "/v1/user/",
			Middleware: []handlerFunc{middleware.NewCustomLogger()},
			Endpoints: []endpoint{
				{fiber.MethodPost, "/add", v1.UserController.AddUser},
			},
		},
	}
	return routes
}

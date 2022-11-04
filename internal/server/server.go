package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-rest-webserver-template/internal/routing"
	"log"
)

type App struct {
	*fiber.App
}

func NewServerApp(ctx context.Context) App {
	//TODO: Make App configurable - JSON serializer
	serverConfig := fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "",
		StrictRouting:                false,
		CaseSensitive:                false,
		Immutable:                    false,
		UnescapePath:                 false,
		ETag:                         false,
		BodyLimit:                    0,
		Concurrency:                  0,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		ReadBufferSize:               0,
		WriteBufferSize:              0,
		CompressedFileSuffix:         "",
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 nil,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        false,
		AppName:                      "",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: false,
		ReduceMemoryUsage:            false,
		JSONEncoder:                  nil,
		JSONDecoder:                  nil,
		XMLEncoder:                   nil,
		Network:                      "",
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnableIPValidation:           false,
		EnablePrintRoutes:            false,
		ColorScheme:                  fiber.Colors{},
	}
	return App{fiber.New(serverConfig)}
}

func (s *App) GetServer() *fiber.App {
	return s.App
}

func (s *App) AddMiddleware(middleware interface{}) {
	s.App.Use(middleware)
}

func (s *App) AddRoute(route routing.Route) *App {
	grpName := route.Group
	routeGrp := s.App.Group(grpName)

	for _, v := range route.Middleware {
		routeGrp.Use(v)
	}

	for _, v := range route.Endpoints {
		switch v.Method {
		case fiber.MethodGet:
			routeGrp.Get(v.Path, v.Handler)
		case fiber.MethodPost:
			routeGrp.Post(v.Path, v.Handler)
		case fiber.MethodPut:
			routeGrp.Put(v.Path, v.Handler)
		case fiber.MethodPatch:
			routeGrp.Patch(v.Path, v.Handler)
		}
	}

	return s
}

// StartServer start rest api server
func (s *App) StartServer(addr string) {
	log.Printf("Starting Listing at %v \n", addr)
	err := s.Listen(":" + addr)
	if err != nil {
		log.Println(err)
		return
	}
}

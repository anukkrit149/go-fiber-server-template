package boot

import (
	"context"
	"go-rest-webserver-template/internal/config"
	"go-rest-webserver-template/internal/database/connection"
	dbservice "go-rest-webserver-template/internal/database/connection/service"
	"go-rest-webserver-template/internal/handler"
	"go-rest-webserver-template/internal/routing"
	"go-rest-webserver-template/internal/server"
	"log"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

type contextKey string

func NewContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	config := config.GetConfig()
	for k, v := range structs.Map(config.Core) {
		key := contextKey(strings.ToLower(k))
		ctx = context.WithValue(ctx, key, v)
	}
	return ctx
}

func InitAPI(ctx context.Context, env string) error {
	return initServerServices(ctx, env)
}

func initServerServices(ctx context.Context, env string) error {
	err := initConfig(env)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//todo: init logger
	//Init DB connection
	dbConn, err := initDatabase(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Init Domains
	initControllers(ctx, dbConn)

	// Service Layer - HTTP Server
	app := initAppServer(ctx)
	// Routes Initialization
	initializeRouter(app)

	log.Println("Initialization complete")
	// Listen HTTP layer
	app.StartServer(strconv.Itoa(config.GetConfig().Core.Port))

	return nil
}

func initAppServer(ctx context.Context) server.App {
	return server.NewServerApp(ctx)
}

func initConfig(env string) error {
	err := config.InitConfig(env)
	return err
}

func initControllers(ctx context.Context, db connection.DBSelector) {
	handler.InitHandlers(ctx, db)
}
func initDatabase(ctx context.Context) (connection.DBSelector, error) {
	db, err := dbservice.NewDBConnection(ctx, config.GetConfig())
	return db, err
}

func initializeRouter(app server.App) {
	routing.Initialize(app.App, config.GetConfig().Core)
}

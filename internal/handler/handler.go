package handler

import (
	"context"
	v1 "go-rest-webserver-template/internal/controller/v1"
	"go-rest-webserver-template/internal/database/connection"
	user "go-rest-webserver-template/internal/domain/user/service"
)

func InitHandlers(ctx context.Context, selector connection.DBSelector) {

	dbConn := selector.GetDBConnection(ctx)
	// Repo initialization
	userRepo, _ := user.NewRepo(ctx, dbConn)
	// Core initialization
	userCore, _ := user.NewUserCore(ctx, userRepo)
	// service logic initialization
	userService := user.NewUserServe(ctx, userCore)
	// controller initialization
	v1.InitUserController(userService)

}

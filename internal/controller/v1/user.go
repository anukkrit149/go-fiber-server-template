package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-webserver-template/internal/domain/user"
)

type IUserController interface {
	AddUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}

var (
	UserController IUserController
)

func InitUserController(service user.IUserServer) {
	UserController = NewUserController(service)
}

func NewUserController(serve user.IUserServer) IUserController {
	return userController{
		serviceLayer: serve,
	}
}

type userController struct {
	serviceLayer user.IUserServer
}

func (c userController) AddUser(ctx *fiber.Ctx) error {
	return nil
}

func (c userController) GetUser(ctx *fiber.Ctx) error {
	return nil
}

func (c userController) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}

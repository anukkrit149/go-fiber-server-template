package service

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/internal/database/model"
	"go-rest-webserver-template/internal/domain/user"
	"go-rest-webserver-template/internal/structs"
)

type UserServe struct {
	issuer user.IUserCore
}

func (u UserServe) Create(ctx context.Context, user *structs.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServe) Get(ctx context.Context, Id uuid.UUID) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserServe(ctx context.Context, issuer user.IUserCore) (user.IUserServe, error) {
	return UserServe{issuer: issuer}, nil
}

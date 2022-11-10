package service

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/internal/database/model"
	"go-rest-webserver-template/internal/domain/user"
	"go-rest-webserver-template/internal/structs"
)

type UserServe struct {
	issuerCore user.IUserCore
}

func (u *UserServe) Create(ctx context.Context, user *structs.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServe) Get(ctx context.Context, id uuid.UUID) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServe) Ops(ctx context.Context, id uuid.UUID) {
	//TODO implement me
	panic("implement me")
}

func NewUserServe(ctx context.Context, issuerCore user.IUserCore) user.IUserServer {
	userServer := UserServe{
		issuerCore: issuerCore,
	}
	return &userServer
}

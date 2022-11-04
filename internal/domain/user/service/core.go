package service

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/internal/database/model"
	"go-rest-webserver-template/internal/domain/user"
	"go-rest-webserver-template/internal/structs"
)

type UserCore struct {
	repo IRepo
}

func (u UserCore) Create(ctx context.Context, user *structs.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserCore) Get(ctx context.Context, Id uuid.UUID) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserCore(ctx context.Context, repo IRepo) (user.IUserCore, error) {
	return UserCore{repo: repo}, nil
}

package user

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/internal/database/model"
	"go-rest-webserver-template/internal/structs"
)

type IUserServe interface {
	Create(ctx context.Context, user *structs.User) (model.User, error)
	Get(ctx context.Context, Id uuid.UUID) (model.User, error)
}

type IUserCore interface {
	Create(ctx context.Context, user *structs.User) (model.User, error)
	Get(ctx context.Context, Id uuid.UUID) (model.User, error)
}

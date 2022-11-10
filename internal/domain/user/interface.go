package user

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/internal/database/model"
	"go-rest-webserver-template/internal/structs"
)

type IUserServer interface {
	Create(ctx context.Context, user *structs.User) (model.User, error)
	Get(ctx context.Context, id uuid.UUID) (*model.User, error)
	Ops(ctx context.Context, id uuid.UUID)
}

type IUserCore interface {
	Create(ctx context.Context, user *structs.User) (model.User, error)
	Get(ctx context.Context, id uuid.UUID) (*model.User, error)
}

package service

import (
	"context"
	"github.com/google/uuid"
	"go-rest-webserver-template/pkg/errors"
	"go-rest-webserver-template/pkg/spine"
	"go-rest-webserver-template/pkg/spine/db"
)

type Repo struct {
	spine.Repo
}

type IRepo interface {
	Create(ctx context.Context, receiver spine.IModel) errors.IError
	FindByID(ctx context.Context, receiver spine.IModel, id uuid.UUID) errors.IError
	FindMany(ctx context.Context, receivers interface{}, condition map[string]interface{}) errors.IError
	Delete(ctx context.Context, receiver spine.IModel) errors.IError
	Update(ctx context.Context, receiver spine.IModel, attrList ...string) errors.IError
}

func NewRepo(ctx context.Context, db *db.DB) (IRepo, error) {
	_ = ctx
	return &Repo{
		Repo: spine.Repo{Db: db},
	}, nil
}

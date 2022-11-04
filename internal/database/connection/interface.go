package connection

import (
	"context"
	"go-rest-webserver-template/pkg/spine/db"
)

type DBSelector interface {
	GetDBConnection(ctx context.Context) *db.DB
}

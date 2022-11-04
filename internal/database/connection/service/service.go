package service

import (
	"context"
	"go-rest-webserver-template/internal/config"
	"go-rest-webserver-template/internal/database/connection"
	"go-rest-webserver-template/pkg/spine/db"
)

var (
	DbConnection *db.DB
)

type DBConnection struct {
	DB *db.DB
}

func NewDBConnection(ctx context.Context, config config.Config) (connection.DBSelector, error) {
	DB, err := db.NewDb(&config.Db)
	DbConnection = DB
	dbConnection := DBConnection{
		DB: DB,
	}

	return dbConnection, err
}

func (d DBConnection) GetDBConnection(ctx context.Context) *db.DB {
	return d.DB
}

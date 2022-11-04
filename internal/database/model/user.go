package model

import (
	"go-rest-webserver-template/pkg/errors"
	"go-rest-webserver-template/pkg/spine"
)

const (
	UserTable = "user"
)

type User struct {
	spine.Model
	UserName string `json:"userName"`
}

func (m *User) SetDefaults() errors.IError {
	//TODO implement me
	panic("implement me")
}

func (m *User) EntityName() string {
	return UserTable
}

func (m *User) TableName() string {
	return UserTable
}

func test(model spine.IModel) {

}

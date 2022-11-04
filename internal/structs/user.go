package structs

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
}

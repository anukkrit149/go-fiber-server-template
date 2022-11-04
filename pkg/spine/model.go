package spine

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"go-rest-webserver-template/pkg/errors"
	"go-rest-webserver-template/pkg/spine/datatype"
)

const (
	AttributeID = "id"
)

type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type IModel interface {
	TableName() string
	EntityName() string
	GetID() uuid.UUID
	Validate() errors.IError
	SetDefaults() errors.IError
}

// Validate validates base Model.
func (m *Model) Validate() errors.IError {
	return GetValidationError(
		validation.ValidateStruct(
			m,
			validation.Field(&m.CreatedAt, validation.By(datatype.IsTimestamp)),
			validation.Field(&m.UpdatedAt, validation.By(datatype.IsTimestamp)),
		),
	)
}

// GetID gets identifier of entity.
func (m *Model) GetID() uuid.UUID {
	return m.ID
}

// GetCreatedAt gets created time of entity
func (m *Model) GetCreatedAt() int64 {
	return m.CreatedAt
}

// GetUpdatedAt gets last updated time of entity
func (m *Model) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// SoftDeletableModel struct is base for entity models
//type SoftDeletableModel struct {
//	Model
//	DeletedAt nulls.Int64 `sql:"DEFAULT:null" json:"deleted_at"`
//}
//
//// Validate validates SoftDeletableOwnedModel.
//func (sm *SoftDeletableModel) Validate() errors.IError {
//	err := GetValidationError(
//		validation.ValidateStruct(
//			sm,
//			validation.Field(&sm.DeletedAt,
//				validation.By(datatype.ValidateNullableInt64(sm.DeletedAt, datatype.IsTimestamp))),
//		),
//	)
//
//	if err == nil {
//		return sm.Model.Validate()
//	}
//
//	return err
//}

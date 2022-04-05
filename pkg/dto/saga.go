package dto

import (
	"orcinator/pkg/domain/saga"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SagaDto struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
	// Members []EventDto
}

func (SagaDto) TableName() string {
	return "sagas"
}

type CreateSagaDto struct {
	SagaDto
}

type RemoveSagaDto struct {
	SagaDto
}

func NewCreateSagaDto(saga saga.Saga) CreateSagaDto {
	dto := CreateSagaDto{}
	dto.ID = saga.Id()
	return dto
}

func NewRemoveSagaDto(saga saga.Saga) RemoveSagaDto {
	dto := RemoveSagaDto{}
	dto.ID = saga.Id()
	return dto
}

func (dto SagaDto) ToSaga() saga.Saga {
	id := uuid.MustParse(dto.ID)
	event := saga.NewEvent("test", saga.NewService("test service"))
	s := saga.New(id, event)
	return s
}

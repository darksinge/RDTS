package dto

import (
	"gorm.io/gorm"
)

type EventDto struct {
	gorm.Model
	ID   string
	Name string
}

// func (dto EventDto) ToEvent() saga.Event {
// 	return saga.NewEvent(name, owner)
// }

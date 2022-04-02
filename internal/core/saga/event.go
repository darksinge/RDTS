package saga

import (
	"github.com/google/uuid"
)

type Event struct {
	Id    uuid.UUID
	Name  string
	Owner *Service
}

func NewEvent(name string, owner Service) Event {
	return Event{
		Id:    uuid.New(),
		Name:  name,
		Owner: &owner,
	}
}

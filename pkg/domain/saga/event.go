package saga

import (
	"github.com/google/uuid"
)

type Event struct {
	Id      uuid.UUID
	Name    string
	Owner   *Service
	schema  string
	version string
}

func NewEvent(name string, owner Service) Event {
	return Event{
		Id:    uuid.New(),
		Name:  name,
		Owner: &owner,
	}
}

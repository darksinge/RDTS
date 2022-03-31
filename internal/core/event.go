package rdts

import "github.com/google/uuid"

type Event struct {
	id    uuid.UUID
	name  string
	owner *Service
}

func NewEvent(name string, owner Service) Event {
	return Event{
		id:    uuid.New(),
		name:  name,
		owner: &owner,
	}
}

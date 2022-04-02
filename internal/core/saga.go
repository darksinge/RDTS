package core

import (
	"github.com/google/uuid"
)

type Saga struct {
	id        string
	members   []*Service
	sequencer sequencer
}

func (s *Saga) Members() []*Service {
	eventList := []*Event{}

	head := &s.sequencer.head
	for head != nil {
		event := head.event
		eventList = append(eventList, event)
		head = head.next
	}

	services := []*Service{}
	svcIds := make(map[uuid.UUID]bool)
	for _, e := range eventList {
		if svcIds[e.Owner.Id] != true {
			services = append(services, e.Owner)
			svcIds[e.Owner.Id] = true
		}
	}

	return services
}

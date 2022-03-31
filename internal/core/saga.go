package rdts

type Saga struct {
	id           string
	members      []*Service
	stateMachine stateMachine
}

func (s *Saga) Members() []*Service {
	events := []*Event{}

	var event *Event
	event = &s.stateMachine.head.event
	for event != nil {
		events = append(events, event)
		event = event
	}

	services := []*Service{}
	for _, e := range events {
		services = append(services, e.owner)
	}
	return services
}

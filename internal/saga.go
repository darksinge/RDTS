package rdts

type Saga struct {
	id           string
	members      []*Service
	stateMachine StateMachine
}

func (s *Saga) Members() []*Service {
	events := []*Event{s.stateMachine.head}
}

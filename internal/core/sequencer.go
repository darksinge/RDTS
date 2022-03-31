package sequencer

import (
	"errors"

	"github.com/google/uuid"
)

//
//
//
//
//
//
// NOTE: Totally just realized this isn't a true state machine because cycles
// aren't allowed. Not much of a state machine!
// TODO: Rename "stateMachine*" to something else? Might need to look into this further.
// Possible alternate names:
// - Sequencer
// - FiniteStateMachine
//
//
// TODO: I should define lifecyle hooks in my state machines
// Possible hooks:
// - beforeTransitionTo or beforeEnter
// - afterTransitionTo or afterEnter
// - beforeAction
// - afterAction
// - beforeTransitionOut or beforeExit
// - afterTransitionOut or afterExit
//
//
//
//
//

type stateMachineNode struct {
	event             Event
	compensatingEvent *Event
	next              *stateMachineNode
}

type stateMachine struct {
	id   uuid.UUID
	head stateMachineNode
	tail *stateMachineNode
}

func checkEventOwners(event Event, compensatingEvent *Event) error {
	// if 'compensatingEvent' is given, then it must be owned by the same Service
	// as 'event'
	if compensatingEvent != nil && event.owner.id == compensatingEvent.owner.id {
		return errors.New("'event' and 'compensatingEvent' must be owned by the same Service.")
	}

	return nil
}

// TODO: Instead of compensatingEvent, maybe it should be an entire state machine
func NewStateMachine(event Event, compensatingEvent *Event) (*stateMachine, error) {
	if err := checkEventOwners(event, compensatingEvent); err != nil {
		return nil, err
	}

	head := stateMachineNode{
		event,
		compensatingEvent,
		nil,
	}

	sm := stateMachine{
		uuid.New(),
		head,
		&head,
	}

	return &sm, nil
}

func (s *stateMachine) AddEvent(event Event, compensatingEvent *Event) error {
	if err := checkEventOwners(event, compensatingEvent); err != nil {
		return err
	}

	// TODO: detect event cycles and return error if found

	node := stateMachineNode{
		event,
		compensatingEvent,
		nil,
	}

	s.tail.next = &node
	s.tail = &node

	return nil
}

func (s *stateMachine) RemoveEvent(event Event) bool {
	node := &s.head
	if node.event.id == event.id {
		s.head = *s.head.next
		return true
	}

	prev := node
	next := node.next
	for prev != nil && next != nil {
		if next.event.id == event.id {
			prev.next = next.next
			return true
		}
	}

	return false
}

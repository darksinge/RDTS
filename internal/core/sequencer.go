package core

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

type sequencerNode struct {
	event             *Event
	compensatingEvent *Event
	next              *sequencerNode
}

type sequencer struct {
	id   uuid.UUID
	head sequencerNode
	tail *sequencerNode
}

func checkEventOwners(event *Event, compensatingEvent *Event) error {
	// if 'compensatingEvent' is given, then it must be owned by the same Service
	// as 'event'
	if compensatingEvent != nil && event.Owner.Id == compensatingEvent.Owner.Id {
		return errors.New("'event' and 'compensatingEvent' must be owned by the same Service.")
	}

	return nil
}

// TODO: Instead of compensatingEvent, maybe it should be an entire state machine
func NewStateMachine(event *Event, compensatingEvent *Event) (*sequencer, error) {
	if err := checkEventOwners(event, compensatingEvent); err != nil {
		return nil, err
	}

	head := sequencerNode{
		event,
		compensatingEvent,
		nil,
	}

	sm := sequencer{
		uuid.New(),
		head,
		&head,
	}

	return &sm, nil
}

func detectCycle(s *sequencer) error {
	visited := make(map[string]bool)

	head := &s.head
	for head != nil {
		e := head.event
		if visited[e.Id.String()] == true {
			return errors.New("detected cycle in sequencer")
		}

		visited[e.Id.String()] = true
		head = head.next
	}

	return nil
}

func (s *sequencer) AddEvent(event *Event, compensatingEvent *Event) error {
	if err := checkEventOwners(event, compensatingEvent); err != nil {
		return err
	}

	node := sequencerNode{
		event,
		compensatingEvent,
		nil,
	}

	s.tail.next = &node
	s.tail = &node

	return detectCycle(s)
}

func (s *sequencer) RemoveEvent(event Event) bool {
	node := &s.head
	if node.event.Id == event.Id {
		s.head = *s.head.next
		return true
	}

	prev := node
	next := node.next
	for prev != nil && next != nil {
		if next.event.Id == event.Id {
			prev.next = next.next
			return true
		}
	}

	return false
}

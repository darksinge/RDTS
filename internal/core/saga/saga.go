package saga

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Saga is the Aggregate Root
//
// A Saga represents a group of events belonging to a distributed transaction.
// A transaction always begins with a single Event, called the initiator.  The
// initiator is always the first event in `members`.
type Saga struct {
	id      uuid.UUID
	members []*Event
}

func New(event Event) Saga {
	members := []*Event{&event}
	return Saga{uuid.New(), members}
}

// Finds the index of `e` in the Saga's event members. Returns -1 if not found.
func (s *Saga) indexOfMember(e Event) int {
	for i, member := range s.members {
		if member.Id == e.Id {
			return i
		}
	}

	return -1
}

// The initiator is defined as the event that triggers a new transaction and is
// always the first element in `s.members`
func (s *Saga) Initiator() *Event {
	return s.members[0]
}

func (s *Saga) Members() []*Event {
	return s.members
}

func (s *Saga) Services() []*Service {
	ids := make(map[uuid.UUID]bool)
	services := []*Service{}
	for _, e := range s.members {
		if !ids[e.Owner.Id] {
			services = append(services, e.Owner)
			ids[e.Owner.Id] = true
		}
	}

	return services
}

func (s *Saga) AddMember(e Event) error {
	if s.indexOfMember(e) > -1 {
		return errors.New("event is already a member of this Saga")
	}

	s.members = append(s.members, &e)
	return nil
}

func (s *Saga) RemoveMember(e Event) (bool, error) {
	index := s.indexOfMember(e)

	if index == -1 {
		return false, nil
	}

	if index == 0 {
		return false, errors.New("cannot remove event from Saga because it is the initiator. Promote a different event to initiator and try again.")
	}

	s.members = append(s.members[:index], s.members[index+1:]...)
	return true, nil
}

func (s *Saga) SetInitiator(e Event) (bool, error) {
	index := s.indexOfMember(e)

	// event isn't a member of this saga
	if index == -1 {
		return false, errors.New(fmt.Sprintf("event must be a member of Saga '%s' to promote to initiator", s.id))
	}

	// event is already the initiator
	if index == 0 {
		return false, nil
	}

	temp := s.members[0]
	s.members[0] = s.members[index]
	s.members[index] = temp

	return true, nil
}

func (s *Saga) StartTransaction() DistributedTransaction {
	return NewDistributedTransaction(*s)
}

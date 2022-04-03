package saga

import (
	"time"

	"github.com/google/uuid"
)

type DistributedTransaction struct {
	id                uuid.UUID
	saga              Saga
	time              time.Time
	deadline          *time.Time
	localTransactions []localTransaction
}

func NewDistributedTransaction(saga Saga) DistributedTransaction {
	trxs := []localTransaction{}
	for _, event := range saga.Members() {
		trxs = append(trxs, newLocalTransaction(*event))
	}

	return DistributedTransaction{
		uuid.New(),
		saga,
		time.Now(),
		nil,
		trxs,
	}
}

func (dtrx *DistributedTransaction) SetDeadline(t time.Time) {
	dtrx.deadline = &t
}

func (dtrx DistributedTransaction) Id() uuid.UUID {
	return dtrx.id
}

func (dtrx DistributedTransaction) Saga() Saga {
	return dtrx.saga
}

func (dtrx *DistributedTransaction) OnDeadlineExceeded() {
	panic("not implemented")
}

func (dtrx *DistributedTransaction) IsComplete() bool {
	for _, t := range dtrx.localTransactions {
		if !t.done {
			return false
		}
	}

	return true
}

// Marks a local transaction as complete.
func (dtrx *DistributedTransaction) EventCompleted(e Event) {
	for _, t := range dtrx.localTransactions {
		if t.event.Id == e.Id {
			t.done = true
			break
		}
	}
}

type localTransaction struct {
	event Event
	time  *time.Time
	done  bool
}

func newLocalTransaction(e Event) localTransaction {
	return localTransaction{e, nil, false}
}

func (dtrx *DistributedTransaction) BeginLocalTransaction(e Event) {
	for _, trx := range dtrx.localTransactions {
		if trx.event.Id == e.Id {
			t := time.Now()
			trx.time = &t
			break
		}
	}
}

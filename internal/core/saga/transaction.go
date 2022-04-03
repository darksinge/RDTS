package saga

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type DistributedTransaction struct {
	id                uuid.UUID
	ctx               *context.Context
	saga              Saga
	time              time.Time
	localTransactions []Transaction
}

func NewDistributedTransaction(ctx *context.Context, saga Saga) DistributedTransaction {
	trxs := []Transaction{}
	for _, event := range saga.Members() {
		trxs = append(trxs, NewTransaction(*event))
	}

	return DistributedTransaction{
		uuid.New(),
		ctx,
		saga,
		time.Now(),
		trxs,
	}
}

func (t DistributedTransaction) Id() uuid.UUID {
	return t.id
}

func (t DistributedTransaction) Saga() Saga {
	return t.saga
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

type Transaction struct {
	event Event
	time  *time.Time
	done  bool
}

func NewTransaction(e Event) Transaction {
	return Transaction{e, nil, false}
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

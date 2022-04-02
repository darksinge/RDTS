package saga

import "github.com/google/uuid"

type Transaction struct {
	id   uuid.UUID
	saga Saga
}

func (t *Transaction) Id() uuid.UUID {
	return t.id
}

func (t *Transaction) Saga() Saga {
	return t.saga
}

package api

import (
	"rdts/internal/core/saga"

	"github.com/google/uuid"
)

type IApi interface {
	RegisterService(name string) (uuid.UUID, error)
	CreateSaga(serviceId string, initiator saga.Event) (saga.Saga, error)
	RegisterEvent(serviceId string, event saga.Event, saga saga.Saga) error
	InitiateTransaction(initiator saga.Event) (bool, error)
	GetServices() ([]saga.Service, error)
	GetServiceEvents(serviceId string) ([]saga.Event, error)
	NotifyDidCompleteTransaction(trxId string, eventId string) error
}

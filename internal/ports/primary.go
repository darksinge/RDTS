package ports

import (
	"github.com/google/uuid"

	"orcinator/internal/domain/saga"
)

type CommandPort interface {
	RegisterService(name string) (uuid.UUID, error)
}

type QueryPort interface {
	GetService(serviceId string) (saga.Saga, error)
}

type ApiPort interface {
	CommandPort
	QueryPort
}

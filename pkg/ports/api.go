package ports

import (
	"github.com/google/uuid"

	"orcinator/internal/domain/saga"
)

type Commands interface {
	RegisterService(name string) (uuid.UUID, error)
}

type Queries interface {
	GetService(serviceId string) (saga.Saga, error)
}

type ApiPort interface {
	Commands
	Queries
}

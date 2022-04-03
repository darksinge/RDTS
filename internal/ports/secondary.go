package ports

import "orcinator/internal/domain/saga"

type DbPort interface {
	CloseConnection() error
	CreateSaga(saga saga.Saga)
	CreateService(svc saga.Service)
}

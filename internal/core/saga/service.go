package saga

import "github.com/google/uuid"

type Service struct {
	Id   uuid.UUID
	Name string
}

func NewService(name string) Service {
	return Service{
		uuid.New(),
		name,
	}
}

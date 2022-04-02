package core

import "github.com/google/uuid"

type Service struct {
	Id   uuid.UUID
	Name string
}

func NewService(id uuid.UUID, name string) Service {
	return Service{
		id,
		name,
	}
}

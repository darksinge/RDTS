package rdts

type Service struct {
	id   string
	name string
}

func NewService(id string, name string) Service {
	return Service{
		id,
		name,
	}
}

package greet

type Service interface {
	Meet(name, gender string) (msg string, err error)
}

type service struct{}

func New() Service {
	return &service{}
}

package math

type Service interface {
	Calculate(operation, a, b string) (msg string, err error)
}

type service struct{}

func New() Service {
	return &service{}
}

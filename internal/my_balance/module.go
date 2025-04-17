package my_balance

type Service interface {
	SomToDiram(somoni string) (diram int, err error)
}

type service struct{}

func New() Service {
	return &service{}
}

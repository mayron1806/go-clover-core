package usecase

type Usecase[Input any, Output any] interface {
	Execute(input Input) (*Output, error)
}

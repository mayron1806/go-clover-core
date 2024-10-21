package usecase

type Usecase interface {
	Execute(input any) (*any, error)
}

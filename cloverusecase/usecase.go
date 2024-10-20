package cloverusecase

type Usecase interface {
	Execute(input any) (*any, error)
}

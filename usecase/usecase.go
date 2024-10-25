package usecase

import "github.com/mayron1806/go-clover-core"

type IUsecase[Input any, Output any] interface {
	Execute(input Input) (*Output, error)
	New(repositories map[string]*clover.Repository) *IUsecase[Input, Output]
}

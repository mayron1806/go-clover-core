package usecase

import "github.com/mayron1806/go-clover-core/repository"

type InputCreate[E any] interface {
	ToEntity() E
}
type UseCaseCreate[E any] struct {
	repository repository.Repository[E]
}

func (u UseCaseCreate[E]) Execute(input InputCreate[E]) (*E, error) {
	entity := input.ToEntity()

	newEntity, err := u.repository.Create(entity)

	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

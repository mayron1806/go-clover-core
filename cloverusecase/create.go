package cloverusecase

import "github.com/mayron1806/go-clover-core/cloverepo"

type InputCreate[E any] interface {
	ToEntity() E
}
type UseCaseCreate[E any] struct {
	repository cloverepo.Repository[E]
}

func (u UseCaseCreate[E]) Execute(input InputCreate[E]) (*E, error) {
	entity := input.ToEntity()

	newEntity, err := u.repository.Create(entity)

	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

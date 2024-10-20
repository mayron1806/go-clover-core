package cloverusecase

import "github.com/mayron1806/go-clover-core/cloverepo"

type UseCaseFind[E any] struct {
	repository cloverepo.Repository[E]
}

func (u UseCaseFind[E]) Execute(id int64) (*E, error) {
	newEntity, err := u.repository.Find(id)

	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

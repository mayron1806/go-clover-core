package usecase

import "github.com/mayron1806/go-fast/repository"

type UseCaseFind[E any] struct {
	repository repository.Repository[E]
}

func (u UseCaseFind[E]) Execute(id int64) (*E, error) {
	newEntity, err := u.repository.Find(id)

	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

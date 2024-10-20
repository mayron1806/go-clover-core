package usecase

import (
	"github.com/mayron1806/go-clover-core/model"
	"github.com/mayron1806/go-clover-core/repository"
)

type InputUpdate[E model.IModel] struct {
	key    int64
	entity E
}
type Update[E model.IModel] struct {
	repository repository.Repository[E]
}

func (u Update[E]) Execute(input InputUpdate[E]) (*E, error) {
	entity, err := u.repository.Find(input.key)
	if err != nil {
		return nil, err
	}
	entity.Update(input.entity)
	newEntity, err := u.repository.Update(entity)

	if err != nil {
		return nil, err
	}
	return &newEntity, nil
}

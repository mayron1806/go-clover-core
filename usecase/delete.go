package usecase

import "github.com/mayron1806/go-clover-core/repository"

type UseCaseDelete[E any] struct {
	repository repository.Repository[E]
}

func (u UseCaseDelete[E]) Execute(id int64) error {
	err := u.repository.Delete(id)
	return err
}

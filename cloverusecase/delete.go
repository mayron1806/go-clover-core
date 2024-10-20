package cloverusecase

import "github.com/mayron1806/go-clover-core/cloverepo"

type UseCaseDelete[E any] struct {
	repository cloverepo.Repository[E]
}

func (u UseCaseDelete[E]) Execute(id int64) error {
	err := u.repository.Delete(id)
	return err
}

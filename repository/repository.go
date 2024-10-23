package repository

import "github.com/mayron1806/go-clover-core/model"

type Repository[T model.BaseModel] interface {
	Find(key int64) (*T, error)
	FindAll() ([]*T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(key int64) error
}

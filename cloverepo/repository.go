package cloverepo

type Repository[Entity any] interface {
	Find(key int64) (Entity, error)
	FindAll() ([]Entity, error)
	Create(entity Entity) (Entity, error)
	Update(entity Entity) (Entity, error)
	Delete(key int64) error
}

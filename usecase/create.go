package usecase

// import "github.com/mayron1806/go-clover-core/repository"

// type UseCaseCreate[E any] struct {
// 	repository repository.Repository[E]
// }

// func (u UseCaseCreate[E]) Execute(input E) (any, error) {
// 	newEntity, err := u.repository.Create(input)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return newEntity, nil
// }

// func NewUseCaseCreate[E any](repository repository.Repository[E]) UseCaseCreate[E] {
// 	return UseCaseCreate[E]{
// 		repository: repository,
// 	}
// }

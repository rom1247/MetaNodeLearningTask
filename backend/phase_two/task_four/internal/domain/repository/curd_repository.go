package repository

// CurdRepository 泛型通用接口
type CurdRepository[T any] interface {
	/*	Create( entity *T) error
		Save( entity *T) error
		Update( entity *T) error
		Delete( entity *T) error
		DeleteById( id uint) error
		FindByID( id uint) (*T, error)
		FindAll(ctx context.Context) ([]T, error)
		Transaction( fn func(repo CurdRepository[T]) error) error*/

	Create(entity *T) error
	Save(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	DeleteById(id uint) error
	FindByID(id uint) (*T, error)
	FindAll() ([]*T, error)
	Transaction(fn func(repo CurdRepository[T]) error) error
}

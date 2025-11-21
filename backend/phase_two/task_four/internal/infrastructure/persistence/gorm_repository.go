package persistence

import (
	"backend/backend/phase_two/task_four/internal/domain/repository"

	"gorm.io/gorm"
)

// Repository GormRepository 通用基础的数据库操作仓储接口 通用设计每个方法是不具备事务的
type Repository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) repository.CurdRepository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *Repository[T]) FindAll() ([]T, error) {
	var list []T
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}
func (r *Repository[T]) Save(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) DeleteById(id uint) error {
	return r.db.Delete(new(T), id).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

// Transaction 支持事务 解释是db的操作不具体事务，即使使用db.Transaction()也是单条sql的单条隐式事务（一般是数据库级别的），所以需要构造一个临时的tx
func (r *Repository[T]) Transaction(fn func(repo repository.CurdRepository[T]) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&Repository[T]{db: tx})
	})
}

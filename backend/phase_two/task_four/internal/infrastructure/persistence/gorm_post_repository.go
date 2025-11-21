package persistence

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"

	"gorm.io/gorm"
)

type GormPostRepository struct {
	repository.CurdRepository[model.Post]
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) repository.PostRepository {
	return &GormPostRepository{
		CurdRepository: NewGormRepository[model.Post](db),
		db:             db,
	}
}

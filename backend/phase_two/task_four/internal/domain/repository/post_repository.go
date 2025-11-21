package repository

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
)

type PostRepository interface {
	CurdRepository[model.Post]
}

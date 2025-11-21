package persistence

import "backend/backend/phase_two/task_four/internal/domain/model"

type GormCommentRepository struct {
	*Repository[model.Comment]
}

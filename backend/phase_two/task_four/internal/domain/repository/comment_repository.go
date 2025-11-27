package repository

import "backend/backend/phase_two/task_four/internal/domain/model"

type CommentRepository interface {
	CurdRepository[model.Comment]

	FindAllByUserIdAndPostId(userId uint, postId uint) ([]*model.Comment, error)
}

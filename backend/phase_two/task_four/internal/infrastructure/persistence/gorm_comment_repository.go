package persistence

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"

	"gorm.io/gorm"
)

type GormCommentRepository struct {
	repository.CurdRepository[model.Comment]
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &GormCommentRepository{NewGormRepository[model.Comment](db), db}
}

func (r *GormCommentRepository) FindAllByUserIdAndPostId(userId uint, postId uint) ([]*model.Comment, error) {
	var comments []*model.Comment
	tx := r.db.Where("user_id = ? AND post_id = ?", userId, postId).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comments, nil
}

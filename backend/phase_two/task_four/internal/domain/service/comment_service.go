package service

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

// NewCommentService 声明构造函数
func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s CommentService) CreateComment(userId uint, postId uint, content string) error {
	return s.repo.Create(&model.Comment{
		UserId:  userId,
		PostID:  postId,
		Content: content,
	})
}

func (s CommentService) FindAllComments(userId uint, postId uint) ([]*model.Comment, error) {
	return s.repo.FindAllByUserIdAndPostId(userId, postId)
}

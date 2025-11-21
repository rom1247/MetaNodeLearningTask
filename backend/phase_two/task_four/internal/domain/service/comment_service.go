package service

import "backend/backend/phase_two/task_four/internal/domain/repository"

type CommentService struct {
	repo repository.CommentRepository
}

// NewCommentService 声明构造函数
func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

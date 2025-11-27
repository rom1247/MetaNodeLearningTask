package app

import (
	"backend/backend/phase_two/task_four/internal/domain/service"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
)

type CommentAppService struct {
	commentService *service.CommentService
}

func NewCommentAppService(commentService *service.CommentService) *CommentAppService {
	return &CommentAppService{}
}

func (s *CommentAppService) CreateComment(userId uint, postId uint, content string) error {
	return s.commentService.CreateComment(userId, postId, content)
}

func (s *CommentAppService) FindAllComments(userId uint, postId uint) ([]response.CommentResponse, error) {
	comments, err := s.commentService.FindAllComments(userId, postId)
	res := make([]response.CommentResponse, len(comments))
	for i, comment := range comments {
		res[i] = response.CommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			UserID:  comment.UserId,
			PostID:  comment.PostID,
		}
	}
	return res, err
}

package service

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"

	"gorm.io/gorm"
)

type PostService struct {
	repo repository.PostRepository
}

func (s *PostService) CreatePost(userId uint, title string, content string) error {
	return s.repo.Create(&model.Post{
		Title:   title,
		Content: content,
		UserID:  userId,
	})
}

func (s *PostService) UpdatePost(id uint, title string, content string) error {
	return s.repo.Update(&model.Post{
		Model:   gorm.Model{ID: id},
		Title:   title,
		Content: content,
	})
}

func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(&model.Post{
		Model: gorm.Model{ID: id},
	})
}

func (s *PostService) FindAllPosts() ([]*model.Post, error) {
	return s.repo.FindAll()
}

func (s *PostService) FindOnePost(id uint) (*model.Post, error) {
	return s.repo.FindByID(id)
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

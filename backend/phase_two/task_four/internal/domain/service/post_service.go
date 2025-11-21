package service

import "backend/backend/phase_two/task_four/internal/domain/repository"

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

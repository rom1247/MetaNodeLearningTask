package service

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(userName string, password string, email string) error {
	user := model.User{
		Username: userName,
		Password: password,
		Email:    email,
	}
	return s.repo.Create(&user)
}

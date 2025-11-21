package service

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"
	"backend/backend/phase_two/task_four/pkg/util"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(userName string, password string, email string) error {
	hash, _ := util.HashPassword(password)
	user := model.User{
		Username: userName,
		Password: hash,
		Email:    email,
	}
	return s.repo.Create(&user)
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) FindByUsername(username string) (*model.User, error) {
	return s.repo.FindByUsername(username)
}

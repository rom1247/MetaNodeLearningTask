package app

import (
	"backend/backend/phase_two/task_four/internal/domain/service"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
)

type UserAppService struct {
	userService *service.UserService
}

func NewUserAppService(userService *service.UserService) *UserAppService {
	return &UserAppService{
		userService: userService,
	}
}

func (s *UserAppService) CreateUser(request request.CreateUserRequest) error {
	return s.userService.CreateUser(request.Username, request.Password, request.Email)
}

package app

import (
	"backend/backend/phase_two/task_four/configs"
	"backend/backend/phase_two/task_four/internal/domain/service"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
	"backend/backend/phase_two/task_four/pkg/util"
	"errors"
)

type LoginAppService struct {
	userService *service.UserService
}

func NewLoginAppService(userService *service.UserService) *LoginAppService {
	return &LoginAppService{
		userService: userService,
	}
}

func (s *LoginAppService) Login(req request.LoginRequest) (*response.LoginResponse, error) {
	user, err := s.userService.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("密码错误")
	}

	//生成token
	jwtConfig := configs.NewJwtConfig()
	token, err := util.GenerateToken(user.ID, jwtConfig.Secret, jwtConfig.ExpiresIn)
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return &response.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Token:    token,
		Expire:   int64(jwtConfig.ExpiresIn.Seconds()),
	}, nil

}

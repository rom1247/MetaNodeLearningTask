package bootstrap

import (
	"backend/backend/phase_two/task_four/internal/domain/service"

	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewPostService,
	service.NewUserService,
	service.NewCommentService,
)

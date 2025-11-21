package bootstrap

import (
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller"

	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	controller.NewPostController,
	controller.NewUserController,
)

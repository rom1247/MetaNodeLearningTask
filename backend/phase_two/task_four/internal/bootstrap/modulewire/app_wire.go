package bootstrap

import (
	"backend/backend/phase_two/task_four/internal/app"

	"github.com/google/wire"
)

var AppServiceSet = wire.NewSet(
	app.NewPostAppService,
	app.NewUserAppService,
	app.NewLoginAppService,
	app.NewCommentAppService,
)

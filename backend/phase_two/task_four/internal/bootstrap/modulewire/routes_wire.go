package bootstrap

import (
	"backend/backend/phase_two/task_four/internal/interfaces/http/routes"

	"github.com/google/wire"
)

var RoutesSet = wire.NewSet(
	routes.NewPostRoutes,
	routes.NewUserRoutes,
	routes.NewLoginRoutes,
)

//go:build wireinject
// +build wireinject

// /+build是1.17的写法，1.18及以后版本使用go:build

package bootstrap

import (
	bootstrap "backend/backend/phase_two/task_four/internal/bootstrap/modulewire"

	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewDB, // GORM DB
		bootstrap.PersistenceSet,
		bootstrap.ServiceSet,
		bootstrap.AppServiceSet,
		bootstrap.ControllerSet,
		bootstrap.RoutesSet,
		// Gin App
		NewApp,
	)

	return nil, nil
}

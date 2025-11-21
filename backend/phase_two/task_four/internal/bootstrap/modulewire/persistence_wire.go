package bootstrap

import (
	"backend/backend/phase_two/task_four/internal/infrastructure/persistence"

	"github.com/google/wire"
)

var PersistenceSet = wire.NewSet(
	persistence.NewGormPostRepository,
	persistence.NewGormUserRepository,
)

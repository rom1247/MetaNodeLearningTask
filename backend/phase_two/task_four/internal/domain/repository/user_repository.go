package repository

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"context"
)

// UserRepository 继承 CurdRepository[User] 并扩展特有方法
type UserRepository interface {
	CurdRepository[model.User]
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

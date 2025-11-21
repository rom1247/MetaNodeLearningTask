package repository

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
)

// UserRepository 继承 CurdRepository[User] 并扩展特有方法
type UserRepository interface {
	CurdRepository[model.User]
	FindByEmail(email string) (*model.User, error)

	FindByUsername(username string) (*model.User, error)
}

package persistence

import (
	"backend/backend/phase_two/task_four/internal/domain/model"
	"backend/backend/phase_two/task_four/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	/*
		组合CurdRepository[User]的接口，这里不能直接写repository.GormUserRepository,
		因为repository.UserRepository本身需要NewGormUserRepository实现，Go 语言不允许结构体嵌入一个还未实现的接口
	*/
	repository.CurdRepository[model.User]
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{NewGormRepository[model.User](db), db}
}

// FindByEmail 实现 User 特有方法
func (r *GormUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

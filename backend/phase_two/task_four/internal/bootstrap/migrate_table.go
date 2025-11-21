package bootstrap

import (
	"backend/backend/phase_two/task_four/configs"
	"backend/backend/phase_two/task_four/internal/domain/model"
	"log"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB, config *configs.AppConfig) {

	if !config.AutoMigrate {
		return
	}

	err := db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		// 如果有更多模型可以继续添加
	)
	if err != nil {
		log.Fatalf("表结构注册失败: %v", err)
	}

	log.Println("表结构注册完成")
}

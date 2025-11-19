package task_three

import (
	"backend/backend/phase_two/task_three/database"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db := database.Init(database.Load(database.SQLite))
	DB = db

}

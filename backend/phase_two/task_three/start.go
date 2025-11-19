package task_three

import (
	"backend/backend/phase_two/task_three/database"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Xdb *sqlx.DB

func Init() {
	db := database.Init(database.Load(database.SQLite))
	DB = db

	xdb := database.InitSqlxDb()
	Xdb = xdb

}

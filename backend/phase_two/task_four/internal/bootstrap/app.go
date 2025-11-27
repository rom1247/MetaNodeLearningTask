package bootstrap

import (
	"backend/backend/phase_two/task_four/configs"
	_ "backend/backend/phase_two/task_four/docs" // swag init 生成的 docs 包
	"backend/backend/phase_two/task_four/internal/infrastructure/cache"
	"backend/backend/phase_two/task_four/internal/infrastructure/middleware"
	"backend/backend/phase_two/task_four/internal/interfaces/http/routes"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	engine      *gin.Engine
	postRoutes  *routes.PostRoutes
	userRoutes  *routes.UserRoutes
	loginRoutes *routes.LoginRoutes
}

func NewApp(db *gorm.DB,
	postRoutes *routes.PostRoutes,
	userRoutes *routes.UserRoutes,
	loginRoutes *routes.LoginRoutes) *App {
	r := gin.Default()

	postRoutes.Register(r)
	userRoutes.Register(r)
	loginRoutes.Register(r)

	return &App{db: db, engine: r, postRoutes: postRoutes, userRoutes: userRoutes, loginRoutes: loginRoutes}
}

func (a *App) Run() {
	//http://localhost:8080/swagger/index.html
	a.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	MigrateTables(a.db, configs.NewConfig())

	cache.InitTokenCache(2*time.Hour, 10000)
	a.engine.Use(middleware.GlobalException(true))
	a.engine.Run(":8080")
}

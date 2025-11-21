package bootstrap

import (
	"backend/backend/phase_two/task_four/configs"
	_ "backend/backend/phase_two/task_four/docs" // swag init 生成的 docs 包
	"backend/backend/phase_two/task_four/internal/interfaces/http/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type App struct {
	db         *gorm.DB
	engine     *gin.Engine
	postRoutes *routes.PostRoutes
	userRoutes *routes.UserRoutes
}

func NewApp(db *gorm.DB, postRoutes *routes.PostRoutes, userRoutes *routes.UserRoutes) *App {
	r := gin.Default()

	postRoutes.Register(r)
	userRoutes.Register(r)
	//rootPath := "./backend/phase_two/task_four"
	//r.GET("/docs/swagger.json", func(c *gin.Context) {
	//	c.File(rootPath + "/docs/swagger.json")
	//})
	// 注册整个 static 目录为静态资源
	//wd, _ := os.Getwd()
	//staticPath := filepath.Join(wd, "backend", "phase_two", "task_four", "static") //绝对路径
	//staticPath := rootPath + "/static" // 你想测试的相对路径
	//r.Static("/static", staticPath)    // 相对 main.go 路径
	//
	//r.GET("/docs", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/static/swagger/index.html")
	//})

	return &App{db: db, engine: r, postRoutes: postRoutes, userRoutes: userRoutes}
}

func (a *App) Run() {
	//http://localhost:8080/swagger/index.html
	a.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	MigrateTables(a.db, configs.LoadConfig())
	a.engine.Run(":8080")
}

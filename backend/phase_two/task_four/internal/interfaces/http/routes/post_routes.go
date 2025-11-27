package routes

import (
	"backend/backend/phase_two/task_four/internal/infrastructure/middleware"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	controller *controller.PostController
}

func NewPostRoutes(controller *controller.PostController) *PostRoutes {
	return &PostRoutes{controller: controller}
}

func (r *PostRoutes) Register(router *gin.Engine) {

	authGroup := router.Group("/api/v1/pub/users/:userId").Use(middleware.JWTAuth())
	{
		authGroup.POST("posts", r.controller.Create)
		authGroup.PUT("posts/:id", r.controller.Update)
		authGroup.DELETE("posts/:id", r.controller.Delete)
	}

	group := router.Group("/api/v1/pub/users/:userId")
	{
		group.GET("posts", r.controller.FindAll)
		group.GET("posts/:id", r.controller.FindOne)
	}
}

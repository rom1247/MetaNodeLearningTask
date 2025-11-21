package routes

import (
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
	group := router.Group("/posts")
	{
		group.POST("", r.controller.Create)
		//group.GET("", r.controller.List)
		//group.GET("/:id", r.controller.Get)
		//group.DELETE("/:id", r.controller.Delete)
	}
}

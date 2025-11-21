package routes

import (
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	controller *controller.UserController
}

func NewUserRoutes(controller *controller.UserController) *UserRoutes {
	return &UserRoutes{
		controller: controller,
	}
}

func (r *UserRoutes) Register(router *gin.Engine) {
	group := router.Group("/api/v1/pri")
	{
		group.POST("/users", r.controller.Create)
		//group.GET("", r.controller.List)
		//group.GET("/:id", r.controller.Get)
		//group.DELETE("/:id", r.controller.Delete)
	}
}

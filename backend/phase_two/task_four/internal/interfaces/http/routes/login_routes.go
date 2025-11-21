package routes

import (
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller"

	"github.com/gin-gonic/gin"
)

type LoginRoutes struct {
	controller *controller.LoginController
}

func NewLoginRoutes(controller *controller.LoginController) *LoginRoutes {
	return &LoginRoutes{
		controller: controller,
	}
}

func (r *LoginRoutes) Register(router *gin.Engine) {
	group := router.Group("/api/v1/pub")
	{
		group.POST("/login", r.controller.Login)
	}
}

package routes

import (
	"backend/backend/phase_two/task_four/internal/infrastructure/middleware"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller"

	"github.com/gin-gonic/gin"
)

type CommentRoutes struct {
	controller *controller.CommentController
}

func NewCommentRoutes(controller *controller.CommentController) *CommentRoutes {
	return &CommentRoutes{
		controller: controller,
	}
}

func (r *CommentRoutes) Register(router *gin.Engine) {

	authGroup := router.Group("/api/v1/pub/users/:userId").Use(middleware.JWTAuth())
	{
		authGroup.POST("/posts/:postId/comments", r.controller.Create)
	}
	group := router.Group("/api/v1/pub/users/:userId")
	{
		group.GET("/posts/:postId/comments", r.controller.FindAll)
	}
}

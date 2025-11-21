package controller

import (
	"backend/backend/phase_two/task_four/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	app *app.PostAppService
}

func NewPostController(app *app.PostAppService) *PostController {
	return &PostController{app: app}
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (c *PostController) Create(ctx *gin.Context) {
	//var request CreatePostRequest
	//if err := ctx.ShouldBindJSON(&request); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//post, err := c.app.CreatePost(request.Title, request.Content)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	ctx.JSON(http.StatusOK, nil)
}

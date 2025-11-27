package controller

import (
	"backend/backend/phase_two/task_four/internal/app"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
	"backend/backend/phase_two/task_four/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	app *app.CommentAppService
}

func NewCommentController(app *app.CommentAppService) *CommentController {
	return &CommentController{
		app: app,
	}
}

// Create 创建评论
// @Summary 创建评论
// @Description 创建评论
// @Tags 评论
// @Accept json
// @Produce json
// @Param userId path uint true "用户ID"
// @Param postId path uint true "文章ID"
// @Param comment body request.CreateCommentRequest true "评论信息"
// @Success 204 {string} string "创建成功"
// @Router /api/v1/pub/users/{userId}/posts/{postId}/comments [post]
func (c *CommentController) Create(ctx *gin.Context) {
	var req request.CreateCommentRequest
	userId, ok := util.GetUintParam(ctx, "userId")
	if !ok {
		return
	}
	postId, ok := util.GetUintParam(ctx, "postId")
	if !ok {
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.app.CreateComment(userId, postId, req.Content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// FindAll 获取用户一篇文章的所有评论
// @Summary 获取用户一篇文章的所有评论
// @Description 获取用户一篇文章的所有评论
// @Tags 评论
// @Accept json
// @Produce json
// @Param userId path uint true "用户ID"
// @Param postId path uint true "文章ID"
// @Success 200 {array} response.CommentResponse "获取成功"
// @Router /api/v1/pub/users/{userId}/posts/{postId}/comments [get]
func (c *CommentController) FindAll(ctx *gin.Context) {
	userId, ok := util.GetUintParam(ctx, "userId")
	if !ok {
		return
	}
	postId, ok := util.GetUintParam(ctx, "postId")
	if !ok {
		return
	}
	comments, err := c.app.FindAllComments(userId, postId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

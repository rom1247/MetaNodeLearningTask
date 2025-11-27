package controller

import (
	"backend/backend/phase_two/task_four/internal/app"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
	"backend/backend/phase_two/task_four/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	app *app.PostAppService
}

func NewPostController(app *app.PostAppService) *PostController {
	return &PostController{app: app}
}

// Create 创建文章
// @Summary 创建文章
// @Description 创建文章
// @Tags 文章
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param userId path uint true "用户ID"
// @Param post body request.CreatePostRequest true "文章信息"
// @Success 204 {string} string "创建成功"
// @Failure 400 {object} response.ErrorResponse "请求参数错误"
// @Router /api/v1/pub/users/{userId}/posts [post]
func (c *PostController) Create(ctx *gin.Context) {

	userId, ok := util.GetUintParam(ctx, "userId")
	if !ok {
		return
	}

	var req request.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.app.CreatePost(userId, req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// Update 更新文章
// @Summary 更新文章
// @Description 更新文章
// @Tags 文章
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path uint true "文章ID"
// @Param userId path uint true "用户ID"
// @Param post body request.CreatePostRequest true "文章信息"
// @Success 204 {string} string "更新成功"
// @Failure 400 {object} response.ErrorResponse "请求参数错误"
// @Router /api/v1/pub/users/{userId}/posts/{id} [put]
func (c *PostController) Update(ctx *gin.Context) {
	id, ok := util.GetUintParam(ctx, "id")
	if !ok {
		return
	}

	var req request.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.app.UpdatePost(id, req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// Delete 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path uint true "文章ID"
// @Param userId path uint true "用户ID"
// @Success 204 {string} string "删除成功"
// @Failure 400 {object} response.ErrorResponse "请求参数错误"
// @Router /api/v1/pub/users/{userId}/posts/{id} [delete]
func (c *PostController) Delete(ctx *gin.Context) {
	id, ok := util.GetUintParam(ctx, "id")
	if !ok {
		return
	}
	err := c.app.DeletePost(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// FindAll 获取所有文章
// @Summary 获取所有文章
// @Description 获取所有文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param userId path uint true "用户ID"
// @Success 200 {array} response.PostResponse "获取成功"
// @Failure 500 {object} response.ErrorResponse "系统异常"
// @Router /api/v1/pub/users/{userId}/posts [get]
func (c *PostController) FindAll(ctx *gin.Context) {
	posts, err := c.app.FindAllPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

// FindOne 获取单个文章
// @Summary 获取单个文章
// @Description 获取单个文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param id path uint true "文章ID"
// @Success 200 {object} response.PostResponse "获取成功"
// @Failure 500 {object} response.ErrorResponse "系统异常"
// @Router /api/v1/pub/users/{userId}/posts/{id} [get]
func (c *PostController) FindOne(ctx *gin.Context) {
	id, ok := util.GetUintParam(ctx, "id")
	if !ok {
		return
	}
	post, err := c.app.FindOnePost(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

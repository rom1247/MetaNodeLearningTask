package controller

import (
	"backend/backend/phase_two/task_four/internal/app"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userAppService *app.UserAppService
}

func NewUserController(userAppService *app.UserAppService) *UserController {
	return &UserController{userAppService}
}

// Create
// @Summary 创建用户
// @Description 创建一个新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "用户信息"
// @Success 204 {string} string "创建成功"
// @Failure 400 {object} response.ErrorResponse "请求参数错误"
// @Failure 500 {object} response.ErrorResponse "系统异常"
// @Router /api/v1/pub/register [post]
func (c *UserController) Create(ctx *gin.Context) {
	var req request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(
			http.StatusBadRequest, "请求参数错误", err.Error()))
		return
	}
	err := c.userAppService.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(
			http.StatusInternalServerError, "系统异常", err.Error()))
		return
	}
	ctx.Status(http.StatusNoContent)
}

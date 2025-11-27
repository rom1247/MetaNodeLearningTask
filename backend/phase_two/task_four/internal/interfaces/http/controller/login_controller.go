package controller

import (
	"backend/backend/phase_two/task_four/internal/app"
	"backend/backend/phase_two/task_four/internal/infrastructure/cache"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/request"
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginAppService *app.LoginAppService
}

func NewLoginController(loginAppService *app.LoginAppService) *LoginController {
	return &LoginController{
		loginAppService: loginAppService,
	}
}

// Login
// @Summary 登录
// @Description 登录
// @Tags 登录
// @Accept json
// @Produce json
// @Param login body request.LoginRequest true "登录信息"
// @Success 200 {object} response.LoginResponse "登录成功"
// @Failure 400 {object} response.ErrorResponse "密码错误"
// @Router /api/v1/pub/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(
			http.StatusBadRequest, "请求参数错误", err.Error()))
		return
	}

	res, err := c.loginAppService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *LoginController) Logout(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.TrimPrefix(auth, "Bearer ")

	cache.DeleteToken(token)

	ctx.JSON(http.StatusOK, gin.H{"message": "登出成功"})
}

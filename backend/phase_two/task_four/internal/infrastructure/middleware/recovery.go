package middleware

import (
	"backend/backend/phase_two/task_four/internal/interfaces/http/controller/response"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GlobalException(includeStack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// panic 统一处理

				// 记录 panic（带 stack trace）
				if includeStack {
					zap.L().Error("[Panic]",
						zap.Any("error", r),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Panic]",
						zap.Any("error", r),
					)
				}
				c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
					500,
					"系统繁忙，请稍后重试",
					fmt.Sprintf("%v", r),
				))
				c.Abort()
				return
			}
		}()
		c.Next()

		// ----- 捕获业务错误 -----
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			// 记录 panic（带 stack trace）
			if includeStack {
				zap.L().Error("[client]",
					zap.Any("error", err),
					zap.String("stack", string(debug.Stack())),
				)
			} else {
				zap.L().Error("[client]",
					zap.Any("error", err),
				)
			}

			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				400,
				"客户端错误",
				err.Error(),
			))
			c.Abort()
			return
		}
	}
}

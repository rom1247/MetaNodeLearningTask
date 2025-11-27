package util

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUintParam 从 ctx.Param 获取指定 URL 参数，并转换成 uint
// 如果参数不存在或格式错误，会返回错误并直接写 400 响应
func GetUintParam(c *gin.Context, paramName string) (uint, bool) {
	paramStr := c.Param(paramName)
	if paramStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s 必须传递", paramName)})
		return 0, false
	}

	value, err := strconv.ParseUint(paramStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s 必须是数字", paramName)})
		return 0, false
	}

	return uint(value), true
}

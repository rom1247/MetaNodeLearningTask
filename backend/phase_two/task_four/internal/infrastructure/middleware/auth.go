package middleware

import (
	"backend/backend/phase_two/task_four/configs"
	"backend/backend/phase_two/task_four/internal/infrastructure/cache"
	"backend/backend/phase_two/task_four/pkg/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. 获取 Authorization Bearer token
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "缺少 Authorization Header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		// 2. 校验 JWT
		claims, err := util.ParseToken(token, configs.NewJwtConfig().Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 无效"})
			c.Abort()
			return
		}

		// 3. 校验 Token 是否在本地缓存（实现主动下线）
		_, ok := cache.GetToken(token)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 已失效或过期"})
			c.Abort()
			return
		}

		userIDStr := strconv.Itoa(int(claims.UserID))

		// 4 自定义 header，例如：X-User-ID: 123
		if h := c.GetHeader("X-User-ID"); h != "" {
			if h != userIDStr {
				c.JSON(http.StatusForbidden, gin.H{"message": "无权限操作其他用户的数据"})
				c.Abort()
				return
			}
		}

		// 4. 校验是否本人（URL path 获取 userId）
		pathUserID := c.Param("userId") // 例如 /users/:userId
		if pathUserID != "" {
			if userIDStr != pathUserID {
				c.JSON(http.StatusForbidden, gin.H{"message": "无权限操作其他用户的数据"})
				c.Abort()
				return
			}
		}

		// 5. 将 userId 注入 context
		c.Set("userId", claims.UserID)

		c.Next()
	}
}

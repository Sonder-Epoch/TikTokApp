package middleware

import (
	"TikTokApp/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 403,
				"status_msg":  "登录体验更多",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}

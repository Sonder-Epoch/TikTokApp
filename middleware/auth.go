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
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "msg": "登录体验更多"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

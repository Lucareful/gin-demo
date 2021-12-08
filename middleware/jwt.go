package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luenci/go-gin-example/models"
)

func LoginJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "请求头中未包含token",
			})
			c.Abort()
			return
		}

		claims, err := models.ParseToken(tokenString)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

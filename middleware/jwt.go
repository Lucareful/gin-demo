package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "请求头中未包含token",
			})
			c.Abort()
			return
		}
		//claims, err := jwt.ParseToken(token)
		//if err != nil {
		//	c.JSON(401, gin.H{
		//		"code":    401,
		//		"message": "token解析失败",
		//	})
		//	c.Abort()
		//	return
		//}
		//c.Set("claims", claims)
		c.Next()
	}
}

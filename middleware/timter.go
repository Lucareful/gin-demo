package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Timter 计算请求耗时 middleware
func Timter() gin.HandlerFunc {
	return func(c *gin.Context) {

		nowTime := time.Now()

		//请求处理
		c.Next()

		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}

// c.Next方法，这个是执行后续中间件请求处理的意思（含没有执行的中间件和我们定义的GET方法处理），这样我们才能获取执行的耗时。
// 也就是在c.Next方法前后分别记录时间，就可以得出耗时。
